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

func (app *Handlers) GetGateByName(params iba.GetIBAGateByNameParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		`select id, name, comment from iba_gates
		where name = $1`,
		params.GateName)

	var rec records.IbaGate
	err = row.Scan(&rec.ID, &rec.Name, &rec.Comment)
	if err == pgx.ErrNoRows {
		return common.NewNotFoundServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("gate not found"),
		})
	}

	if err != nil {
		app.logger.Error("Unable to SELECT", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to SELECT: %s", err.Error()),
		})
	}

	result := models.IBAGate{
		ID:      &rec.ID,
		Name:    &rec.Name,
		Comment: &rec.Comment.String,
	}

	return iba.NewGetIBAGateByNameOK().WithPayload(&result)
}

func createOrUpdateGate(conn *pgxpool.Conn, id *int64, name, comment *string) (*models.IBAGate, error) {

	inserting := false
	// select

	row := conn.QueryRow(context.Background(),
		`select id, name, comment from iba_gates
		where name = $1`,
		name)

	var rec records.IbaGate
	err := row.Scan(&rec.ID, &rec.Name, &rec.Comment)
	if err == pgx.ErrNoRows {
		inserting = true
	}

	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("unable to SELECT: %w", err)
	}

	if inserting {
		row := conn.QueryRow(context.Background(),
			"INSERT INTO iba_gates (name, comment) VALUES ($1, $2) RETURNING id",
			name, comment)

		var id int64
		err = row.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("unable to INSERT: %w", err)
		}

		result := &models.IBAGate{
			ID:      &id,
			Name:    name,
			Comment: comment,
		}

		return result, nil
	}

	// updating
	if id != nil {
		return nil, fmt.Errorf("id can only be updated to DEFAULT: %w", common.ErrBadRequest)
	}

	var sqlParams []string
	sqlArgs := []interface{}{*name}

	if comment != nil {
		sqlArgs = append(sqlArgs, *comment)
		sqlParams = append(sqlParams, fmt.Sprintf("comment = $%d", len(sqlArgs)))
	}

	if len(sqlArgs) == 1 {
		result := &models.IBAGate{
			ID:      &rec.ID,
			Name:    name,
			Comment: comment,
		}

		return result, nil
	}

	row = conn.QueryRow(context.Background(),
		fmt.Sprintf("UPDATE iba_gates SET %s WHERE name = $1 RETURNING id, name, comment", strings.Join(sqlParams, ",")),
		sqlArgs...)

	err = row.Scan(&rec.ID, &rec.Name, &rec.Comment)

	if err != nil {
		return nil, fmt.Errorf("unable to UPDATE: %w", err)
	}

	result := &models.IBAGate{
		ID:      &rec.ID,
		Name:    name,
		Comment: comment,
	}

	return result, nil
}

func (app *Handlers) CreateOrUpdateGate(params iba.CreateOrUpdateIBAGateParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	result, err := createOrUpdateGate(conn, params.Body.ID, params.Body.Name, params.Body.Comment)

	if err != nil {
		app.logger.Error("", zap.Error(err))
		if errors.Is(err, common.ErrBadRequest) {
			return iba.NewCreateOrUpdateIBAGateBadRequest().WithPayload(&models.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		if errors.Is(err, common.ErrNotFound) {
			return iba.NewCreateOrUpdateIBAGateNotFound().WithPayload(&models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		}
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return iba.NewCreateOrUpdateIBAGateOK().WithPayload(result)
}

func (app *Handlers) DeleteGateByName(params iba.DeleteIBAGateByNameParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return common.NewInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM iba_gates WHERE name = $1", params.GateName)
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

	return iba.NewDeleteIBAGateByNameOK()
}
