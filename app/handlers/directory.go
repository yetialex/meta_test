package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/yetialex/meta_test/gen/records"
	"github.com/yetialex/meta_test/gen/swagger/models"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/core"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

func (app *CoreHandlers) CreateDirectory(params core.CreateDirectoryParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO Directories (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at",
		params.Body.Name, params.Body.Description)

	var id int64
	var createdAt, updatedAt strfmt.DateTime
	err = row.Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		app.logger.Error("Unable to INSERT", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to INSERT: %s", err.Error()),
		})
	}

	result := models.DirectoryObject{
		ID:          &id,
		Name:        params.Body.Name,
		Description: params.Body.Description,
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
	}

	return core.NewCreateDirectoryOK().WithPayload(&result)
}

func (app *CoreHandlers) DeleteDirectory(params core.DeleteDirectoryParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM Directories WHERE id = $1", params.DirectoryID)
	if err != nil {
		app.logger.Error("Unable to DELETE", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to DELETE: %s", err.Error()),
		})
	}

	if ct.RowsAffected() == 0 {
		return NewCommonNotFoundServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Unable to DELETE"),
		})
	}

	return core.NewDeleteDirectoryOK()
}

func (app *CoreHandlers) GetDirectory(params core.GetDirectoryParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"SELECT id, name, description, created_at, updated_at, acl FROM Directories WHERE id = $1",
		params.DirectoryID)

	var rec records.Directory
	err = row.Scan(&rec.ID, &rec.Name, &rec.Description, &rec.CreatedAt, &rec.UpdatedAt, &rec.ACL)
	if err == pgx.ErrNoRows {
		return NewCommonNotFoundServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("directory not found"),
		})
	}

	if err != nil {
		app.logger.Error("Unable to SELECT", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to SELECT: %s", err.Error()),
		})
	}

	result := models.DirectoryObject{
		ID:          &rec.ID,
		Name:        &rec.Name,
		Description: rec.Description,
		CreatedAt:   &rec.CreatedAt,
		UpdatedAt:   &rec.UpdatedAt,
	}

	return core.NewGetDirectoryOK().WithPayload(&result)
}

func (app *CoreHandlers) UpdateDirectory(params core.UpdateDirectoryParams) middleware.Responder {
	conn, err := app.pool.Acquire(context.Background())
	if err != nil {
		app.logger.Error("Unable to acquire a database connection", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to acquire a database connection: %s", err.Error()),
		})
	}
	defer conn.Release()

	var sqlParams []string
	sqlArgs := []interface{}{params.DirectoryID}

	if params.Body.Name != nil {
		sqlArgs = append(sqlArgs, params.Body.Name)
		sqlParams = append(sqlParams, fmt.Sprintf("name = $%d", len(sqlArgs)))
	}

	if params.Body.Description != nil {
		sqlArgs = append(sqlArgs, params.Body.Description)
		sqlParams = append(sqlParams, fmt.Sprintf("description = $%d", len(sqlArgs)))
	}

	if len(sqlArgs) == 1 {
		return core.NewUpdateDirectoryBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("bad request, no changes"),
		})
	}

	row := conn.QueryRow(context.Background(),
		fmt.Sprintf("UPDATE Directories SET %s WHERE id = $1 RETURNING id, name, description, created_at, updated_at, acl", strings.Join(sqlParams, ",")),
		sqlArgs...)

	var rec records.Directory
	err = row.Scan(&rec.ID, &rec.Name, &rec.Description, &rec.CreatedAt, &rec.UpdatedAt, &rec.ACL)
	if err == pgx.ErrNoRows {
		return NewCommonNotFoundServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("directory not found"),
		})
	}

	if err != nil {
		app.logger.Error("Unable to UPDATE", zap.Error(err))
		return NewCommonInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to UPDATE: %s", err.Error()),
		})
	}

	result := models.DirectoryObject{
		ID:          &rec.ID,
		Name:        &rec.Name,
		Description: rec.Description,
		CreatedAt:   &rec.CreatedAt,
		UpdatedAt:   &rec.UpdatedAt,
	}

	return core.NewUpdateDirectoryOK().WithPayload(&result)
}
