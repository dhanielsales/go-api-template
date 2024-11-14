package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
)

// PaginationResult represents pagination parameters with limit and offset values.
type PaginationResult struct {
	Limit  int32
	Offset int32
}

const (
	DESC = "DESC" // Descending order.
	ASC  = "ASC"  // Ascending order.
)

// Pagination calculates and returns the pagination limit and offset based on the page and perPage parameters.
func Pagination(page, perPage string) PaginationResult {
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		currentPage = 1
	}

	currentPerPage, err := strconv.Atoi(perPage)
	if err != nil {
		currentPerPage = 10
	}

	var limit int
	var offset int

	if currentPage == 0 {
		currentPage = 1
	}

	if currentPerPage == 0 {
		currentPerPage = 10
	}

	if currentPerPage > 100 {
		currentPerPage = 100
	}

	limit = currentPerPage
	offset = (currentPage - 1) * currentPerPage

	//nolint:gosec // safe conversion int to int32
	return PaginationResult{Limit: int32(limit), Offset: int32(offset)}
}

// Sorting constructs an SQL sorting clause for the specified field and direction.
func Sorting(field, direction string) string {
	var orderBy string

	if field == "" {
		field = "created_at"
	}

	if direction == "" {
		direction = DESC
	}

	if direction != ASC && direction != DESC {
		direction = DESC
	}

	orderBy = field + " " + direction

	return orderBy
}

// CallTx manages a database transaction, handling commit or rollback based on the function's success.
func CallTx[R any](ctx context.Context, db *sql.DB, f func(q *sql.Tx) (R, error)) (R, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return *new(R), err
	}

	result, err := f(tx)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return *new(R), errors.Join(err, errTx)
		}
		return *new(R), err
	}

	err = tx.Commit()
	if err != nil {
		return *new(R), err
	}

	return result, nil
}
