package postgres

import (
	"context"
	"database/sql"
	"errors"
)

type PaginationResult struct {
	Limit  int32
	Offset int32
}

func Pagination(page, perPage int32) PaginationResult {
	var limit int32
	var offset int32

	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = 10
	}

	if perPage > 100 {
		perPage = 100
	}

	limit = perPage
	offset = (page - 1) * perPage

	return PaginationResult{
		Limit:  limit,
		Offset: offset,
	}
}

func Sorting(field, direction string) string {
	var orderBy string

	if field == "" {
		field = "created_at"
	}

	if direction == "" {
		direction = "DESC"
	}

	if direction != "ASC" && direction != "DESC" {
		direction = "DESC"
	}

	orderBy = field + " " + direction

	return orderBy
}

func CallTx[R any](ctx context.Context, db *sql.DB, f func(q *sql.Tx) (*R, error)) (*R, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result, err := f(tx)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return nil, errors.Join(err, errTx)
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
