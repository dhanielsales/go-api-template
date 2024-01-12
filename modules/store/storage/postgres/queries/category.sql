-- name: CreateCategory :execresult
INSERT
	INTO category (id, name, slug, description, created_at)
	VALUES ($1, $2, $3, $4, $5);

-- name: UpdateCategory :execresult
UPDATE category
	SET 
		name = $1,
		slug = $2,
		description = $3,
		updated_at = $4
	WHERE id = $5;

-- name: DeleteCategory :execresult
DELETE
	FROM category
	WHERE id = $1;

-- name: GetCategoryById :one
SELECT *
	FROM category
	WHERE id = $1;

-- name: GetManyCategory :many
SELECT *
	FROM category
	ORDER BY sqlc.arg('orderBy')::text
	LIMIT sqlc.arg('limit')::int
	OFFSET sqlc.arg('offset')::int;
