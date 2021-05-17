package iba

import (
	"context"
	"fmt"
	"strings"

	"github.com/yetialex/meta_test/app/handlers/common"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yetialex/meta_test/gen/records"
)

const (
	IbaBatchMode = 2
)

func createOrUpdateSource(conn *pgxpool.Conn, id, parentID *int64, name, comment *string, sourceClassesID int) (*int64, error) {

	inserting := false
	// select

	row := conn.QueryRow(context.Background(),
		`select id, name from sources
		where name = $1`,
		name)

	var rec records.Source
	err := row.Scan(&rec.ID, &rec.Name)
	if err == pgx.ErrNoRows {
		inserting = true
	}

	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("unable to SELECT: %w", err)
	}

	if inserting {
		row := conn.QueryRow(context.Background(),
			"INSERT INTO sources (parent_id, name, source_classes_id, comment) VALUES ($1, $2, $3, $4) RETURNING id",
			parentID, name, sourceClassesID, comment)

		var sourceID int64
		err = row.Scan(&sourceID)
		if err != nil {
			return nil, fmt.Errorf("unable to INSERT: %w", err)
		}

		return &sourceID, nil
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

	if parentID != nil {
		sqlArgs = append(sqlArgs, *parentID)
		sqlParams = append(sqlParams, fmt.Sprintf("parent_id = $%d", len(sqlArgs)))
	}

	if len(sqlArgs) == 1 {
		return nil, fmt.Errorf("no changes: %w", common.ErrBadRequest)
	}

	row = conn.QueryRow(context.Background(),
		fmt.Sprintf("UPDATE sources SET %s WHERE name = $1 RETURNING id, name, comment", strings.Join(sqlParams, ",")),
		sqlArgs...)

	err = row.Scan(&rec.ID, &rec.Name, &rec.Comment)

	if err != nil {
		return nil, fmt.Errorf("unable to UPDATE: %w", err)
	}

	return &rec.ID, nil
}
