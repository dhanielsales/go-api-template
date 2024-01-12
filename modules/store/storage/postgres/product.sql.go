// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createProduct = `-- name: CreateProduct :execresult
INSERT
	INTO product (id, name, slug, description, price, category_id, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateProductParams struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	Description sql.NullString
	Price       float64
	CategoryID  uuid.UUID
	CreatedAt   int64
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProduct,
		arg.ID,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.Price,
		arg.CategoryID,
		arg.CreatedAt,
	)
}

const deleteProduct = `-- name: DeleteProduct :execresult
DELETE
	FROM product
	WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteProduct, id)
}

const getManyProduct = `-- name: GetManyProduct :many
SELECT id, name, slug, description, price, category_id, created_at, updated_at
	FROM product
	ORDER BY $1::text
	LIMIT $3::int
	OFFSET $2::int
`

type GetManyProductParams struct {
	OrderBy string
	Offset  int32
	Limit   int32
}

func (q *Queries) GetManyProduct(ctx context.Context, arg GetManyProductParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getManyProduct, arg.OrderBy, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.Description,
			&i.Price,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getManyProductByCategoryId = `-- name: GetManyProductByCategoryId :many
SELECT id, name, slug, description, price, category_id, created_at, updated_at
	FROM product
	WHERE category_id = $1
`

func (q *Queries) GetManyProductByCategoryId(ctx context.Context, categoryID uuid.UUID) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getManyProductByCategoryId, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.Description,
			&i.Price,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductById = `-- name: GetProductById :one
SELECT id, name, slug, description, price, category_id, created_at, updated_at
	FROM product
	WHERE id = $1
`

func (q *Queries) GetProductById(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.Price,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :execresult
UPDATE product
	SET 
		name = $1,
		slug = $2,
		description = $3,
		price = $4,
		category_id = $5,
		updated_at = $6
	WHERE id = $7
`

type UpdateProductParams struct {
	Name        string
	Slug        string
	Description sql.NullString
	Price       float64
	CategoryID  uuid.UUID
	UpdatedAt   sql.NullInt64
	ID          uuid.UUID
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProduct,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.Price,
		arg.CategoryID,
		arg.UpdatedAt,
		arg.ID,
	)
}
