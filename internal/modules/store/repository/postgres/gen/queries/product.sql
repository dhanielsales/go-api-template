-- name: CreateProduct :execresult
INSERT
	INTO product (id, name, slug, description, price, category_id, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateProduct :execresult
UPDATE product
	SET 
		name = $1,
		slug = $2,
		description = $3,
		price = $4,
		category_id = $5,
		updated_at = $6
	WHERE id = $7;

-- name: DeleteProduct :execresult
DELETE
	FROM product
	WHERE id = $1;

-- name: GetProductById :one
SELECT *
	FROM product
	WHERE id = $1;

-- name: GetManyProduct :many
SELECT *
	FROM product
	ORDER BY sqlc.arg('orderBy')::text
	LIMIT sqlc.arg('limit')::int
	OFFSET sqlc.arg('offset')::int;

-- name: GetManyProductByCategoryId :many
SELECT *
	FROM product
	WHERE category_id = $1;
