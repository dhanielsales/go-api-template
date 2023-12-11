package category_storage

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/golang-scaffold/internal/time"

	category "github.com/dhanielsales/golang-scaffold/modules/category/entity"
)

// This is how the category is stored in postgres
type categorySchema struct {
	Id          string         `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	CreatedAt   int64          `db:"created_at"`
	UpdatedAt   int64          `db:"updated_at"`
}

func (s *CategoryStorage) Create(ctx context.Context, data *category.Category) error {
	_, err := s.postgresDb.Client.ExecContext(ctx, "INSERT INTO category (id, name, description, created_at) VALUES ($1, $2, $3, $4)", data.Id, data.Name, data.Description, data.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryStorage) Update(ctx context.Context, id string, data *category.Category) (int64, error) {
	updatedAt := time.Now()

	// TODO Falar com o Murilo, o que ele prefere, responsabilidade do Repository de fazer o update ou do Service? Ou até mesmo da propria Entity?
	result, err := s.postgresDb.Client.ExecContext(ctx, `
		UPDATE category 
			SET 
				name = $1,
				description = $2,
				updated_at = $3,
		WHERE id = $4
	`, data.Name, data.Description, updatedAt, data.Id)

	if err != nil {
		return 0, err
	}

	affects, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affects, nil
}

func (s *CategoryStorage) GetAll(ctx context.Context) (*[]categorySchema, error) {
	var result []categorySchema
	err := s.postgresDb.Client.SelectContext(ctx, &result, "SELECT * FROM category")
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CategoryStorage) GetById(ctx context.Context, id string) (*categorySchema, error) {
	var result categorySchema
	err := s.postgresDb.Client.GetContext(ctx, &result, "SELECT * FROM category WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CategoryStorage) Delete(ctx context.Context, id string) (int64, error) {
	result, err := s.postgresDb.Client.ExecContext(ctx, "DELETE FROM category WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	affects, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affects, nil
}
