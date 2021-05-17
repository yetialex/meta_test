package iba

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"

	"github.com/yetialex/meta_test/app/handlers/common"

	"github.com/yetialex/meta_test/gen/records"
	"github.com/yetialex/meta_test/gen/swagger/models"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/iba"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func getGateSourceByGateNameAndMnt(conn *pgxpool.Conn, name, mnt string) (*models.IBAGateSource, error) {
	row := conn.QueryRow(context.Background(),
		`select source_id, gate_id, mnt, comment from iba_gates_sources
		where gate_id in (select id from iba_gates where name = $1)
			and mnt = $2`,
		name, mnt)

	var rec records.IbaGatesSource
	err := row.Scan(&rec.SourceID, &rec.GateID, &rec.Mnt, &rec.Comment)
	if err != nil {
		return nil, err
	}

	result := &models.IBAGateSource{
		Comment:     rec.Comment,
		IbaGateID:   &rec.GateID,
		IbaSourceID: &rec.SourceID,
		Mnt:         &rec.Mnt,
	}
	return result, nil
}

func (app *Handlers) GetGateSourceByGateNameAndMnt(params iba.GetGateSourceByGateNameAndMntParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	result, err := getGateSourceByGateNameAndMnt(conn, params.GateName, params.Mnt)

	if errors.Is(err, pgx.ErrNoRows) {
		return common.NewNotFoundServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("gate source not found"),
		})
	}

	if err != nil {
		app.logger.Error("Unable to SELECT", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to SELECT: %s", err.Error()),
		})
	}

	return iba.NewGetGateSourceByGateNameAndMntOK().WithPayload(result)
}

func createOrUpdateGateSourceByGateNameAndMnt(conn *pgxpool.Conn, gateName, mnt, comment *string) (*models.IBAGateSource, error) {

	txOpts := pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	}
	tx, err := conn.BeginTx(context.Background(), txOpts)
	if err != nil {
		return nil, fmt.Errorf("unable to begin serializable transaction: %w", err)
	}
	// Rollback has no effect if Commit will be called
	defer tx.Rollback(context.Background())

	inserting := false
	// select

	row := conn.QueryRow(context.Background(),
		`select source_id, gate_id, mnt, comment from iba_gates_sources
		where gate_id in (select id from iba_gates where name = $1) and mnt = $2`,
		gateName, mnt)

	var rec records.IbaGatesSource
	err = row.Scan(&rec.SourceID, &rec.GateID, &rec.Mnt, &rec.Comment)
	if errors.Is(err, pgx.ErrNoRows) {
		inserting = true
	}

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("unable to SELECT: %w", err)
	}

	if inserting {

		gate, err := createOrUpdateGate(conn, nil, gateName, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to modify gate: %w", err)
		}

		gateID := *gate.ID

		sourceName := fmt.Sprintf("gate=%s/mnt=%s", *gateName, *mnt)
		sourceID, err := createOrUpdateSource(conn, nil, nil, &sourceName, nil, IbaBatchMode)
		if err != nil {
			return nil, fmt.Errorf("unable to modify gate: %w", err)
		}

		row := conn.QueryRow(context.Background(),
			"INSERT INTO iba_gates_sources (source_id, gate_id, mnt, comment) VALUES ($1, $2, $3, $4) RETURNING source_id",
			sourceID, gateID, mnt, comment)

		var id int64
		err = row.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("unable to INSERT: %w", err)
		}

		result := &models.IBAGateSource{
			Comment:     *comment,
			IbaGateID:   &gateID,
			IbaSourceID: sourceID,
			Mnt:         mnt,
		}

		err = tx.Commit(context.Background())
		if err != nil {
			return nil, fmt.Errorf("unable to COMMIT: %w", err)
		}

		return result, nil
	}

	// updating

	var sqlParams []string
	sqlArgs := []interface{}{gateName}

	if comment != nil {
		sqlArgs = append(sqlArgs, comment)
		sqlParams = append(sqlParams, fmt.Sprintf("comment = $%d", len(sqlArgs)))
	}

	if len(sqlArgs) == 1 {
		result := &models.IBAGateSource{
			Comment:     rec.Comment,
			IbaGateID:   &rec.GateID,
			IbaSourceID: &rec.SourceID,
			Mnt:         &rec.Mnt,
		}

		err = tx.Commit(context.Background())
		if err != nil {
			return nil, fmt.Errorf("unable to COMMIT: %w", err)
		}
		return result, nil
	}

	row = conn.QueryRow(context.Background(),
		fmt.Sprintf("UPDATE iba_gates_sources SET %s WHERE gate_id in (select id from iba_gates where name = $1) RETURNING source_id, gate_id, mnt, comment", strings.Join(sqlParams, ",")),
		sqlArgs...)

	err = row.Scan(&rec.SourceID, &rec.GateID, &rec.Mnt, &rec.Comment)

	if err != nil {
		return nil, fmt.Errorf("unable to UPDATE: %w", err)
	}

	result := &models.IBAGateSource{
		Comment:     rec.Comment,
		IbaGateID:   &rec.GateID,
		IbaSourceID: &rec.SourceID,
		Mnt:         &rec.Mnt,
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to COMMIT: %w", err)
	}

	return result, nil

}

func (app *Handlers) CreateOrUpdateGateSourceByGateNameAndMnt(params iba.CreateOrUpdateGateSourceByGateNameAndMntParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	result, err := createOrUpdateGateSourceByGateNameAndMnt(conn, &params.GateName, &params.Mnt, &params.Body.Comment)

	if err != nil {
		app.logger.Error("", zap.Error(err))
		if errors.Is(err, common.ErrBadRequest) {
			return iba.NewCreateOrUpdateGateSourceByGateNameAndMntBadRequest().WithPayload(&models.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		if errors.Is(err, common.ErrNotFound) {
			return iba.NewCreateOrUpdateGateSourceByGateNameAndMntNotFound().WithPayload(&models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		}
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return iba.NewCreateOrUpdateGateSourceByGateNameAndMntOK().WithPayload(result)
}

//
//func (app *Handlers) GetGateSourceByGateNameAndMnt(params iba.GetGateSourceByGateNameAndMntParams) middleware.Responder {
//	conn, err := app.pool.Acquire(context.Background())
//	if err != nil {
//		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
//		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
//			Code:    http.StatusInternalServerError,
//			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
//		})
//	}
//	defer conn.Release()
//
//	rows, err := conn.Query(context.Background(),
//		`select source_id, gate_id, mnt, comment from iba_gates_sources
//		where gate_id in (select id from iba_gates where name = $1)
//			and mnt = '$2'`,
//		params.GateName, params.Mnt)
//	defer rows.Close()
//
//	if err != nil {
//		app.logger.Error("Unable to SELECT", zap.Error(err))
//		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
//			Code:    http.StatusInternalServerError,
//			Message: fmt.Sprintf("Unable to SELECT: %s", err.Error()),
//		})
//	}
//
//	var results iba.GetIBAGateMntsOKBody
//
//	for rows.Next() {
//		var rec records.IbaGatesMnt
//		err = rows.Scan(&rec.SourceID, &rec.GateID, &rec.Mnt, &rec.Comment)
//		if err != nil {
//			app.logger.Error("rows.Scan failed", zap.Error(err))
//			return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
//				Code:    http.StatusInternalServerError,
//				Message: fmt.Sprintf("rows.Scan failed: %s", err.Error()),
//			})
//		}
//		results.Items = append(results.Items, &models.IBAGateMnt{
//			Comment:     rec.Comment,
//			IbaGateID:   &rec.GateID,
//			IbaSourceID: &rec.SourceID,
//			Mnt:         &rec.Mnt,
//		})
//	}
//
//	if rows.Err() != nil {
//		app.logger.Error("conn.Query failed", zap.Error(rows.Err()))
//		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
//			Code:    http.StatusInternalServerError,
//			Message: fmt.Sprintf("conn.Query failed: %s", rows.Err().Error()),
//		})
//	}
//
//	return iba.NewGetIBAGateMntsOK().WithPayload(&results)
//}

func (app *Handlers) DeleteGateSourceByGateNameAndMnt(params iba.DeleteGateSourceByGateNameAndMntParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM iba_gates_sources WHERE gate_id in (select id from iba_gates where name = $1)", params.GateName)
	if err != nil {
		app.logger.Error("Unable to DELETE", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to DELETE: %s", err.Error()),
		})
	}

	if ct.RowsAffected() == 0 {
		return common.NewNotFoundServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Unable to DELETE"),
		})
	}

	return iba.NewDeleteGateSourceByGateNameAndMntOK()
}
