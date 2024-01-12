package application

import (
	"database/sql"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/modules/store/entity"
	"github.com/dhanielsales/golang-scaffold/modules/store/storage"
	store_storage "github.com/dhanielsales/golang-scaffold/modules/store/storage/postgres"
)

type CreateCategoryPayload struct {
	Name        string
	Description string
}

func (s *StoreService) CreateCategory(ctx context.Context, data CreateCategoryPayload) (*int64, error) {
	category := entity.NewCategory(data.Name, data.Description)

	dbResult, err := s.storage.Queries.CreateCategory(ctx, store_storage.CreateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: sql.NullString{String: *category.Description},
		CreatedAt:   category.CreatedAt,
	})

	if err != nil {
		return nil, err
	}

	affected, err := dbResult.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affected, nil
}

func (s *StoreService) GetCategoryById(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	return postgres.CallTx(ctx, s.storage.Db.Client, func(tx *sql.Tx) (*entity.Category, error) {
		queries := s.storage.Queries.WithTx(tx)

		dbResult, err := queries.GetCategoryById(ctx, id)

		if err != nil {
			return nil, err
		}

		dbResultProd, err := queries.GetManyProductByCategoryId(ctx, id)
		if err != nil {
			return nil, err
		}

		var products []entity.Product = []entity.Product{}

		for _, prod := range dbResultProd {
			curr := storage.ToProduct(&prod)
			products = append(products, *curr)
		}

		res := storage.ToCategory(&dbResult)
		res.Products = &products

		return res, nil
	})
}

type GetManyCategoryParams struct {
	Page           string
	PerPage        string
	OrderBy        string
	OrderDirection string
}

func (s *StoreService) GetManyCategory(ctx context.Context, params GetManyCategoryParams) (*[]entity.Category, error) {
	return postgres.CallTx(ctx, s.storage.Db.Client, func(tx *sql.Tx) (*[]entity.Category, error) {
		queries := s.storage.Queries.WithTx(tx)

		pagination := postgres.Pagination(params.Page, params.PerPage)
		sorting := postgres.Sorting(params.OrderBy, params.OrderDirection)

		dbResult, err := queries.GetManyCategory(ctx, store_storage.GetManyCategoryParams{
			Limit:   pagination.Limit,
			Offset:  pagination.Offset,
			OrderBy: sorting,
		})

		if err != nil {
			return nil, err
		}

		var result []entity.Category = []entity.Category{}

		for _, dbCategory := range dbResult {
			curr := storage.ToCategory(&dbCategory)
			result = append(result, *curr)
		}

		return &result, nil
	})
}

type UpdateCategoryPayload struct {
	Name        string
	Description string
}

func (s *StoreService) UpdateCategory(ctx context.Context, id uuid.UUID, data UpdateCategoryPayload) (*int64, error) {
	return postgres.CallTx(ctx, s.storage.Db.Client, func(tx *sql.Tx) (*int64, error) {
		queries := s.storage.Queries.WithTx(tx)

		res, err := queries.GetCategoryById(ctx, id)
		if err != nil {
			return nil, err
		}

		category := storage.ToCategory(&res)

		category.Update(data.Name, data.Description)

		dbResult, err := queries.UpdateCategory(ctx, store_storage.UpdateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Slug:        category.Slug,
			Description: sql.NullString{String: data.Description},
			UpdatedAt:   sql.NullInt64{Int64: *category.UpdatedAt, Valid: true},
		})
		if err != nil {
			return nil, err
		}

		affected, err := dbResult.RowsAffected()
		if err != nil {
			return nil, err
		}

		return &affected, nil
	})
}

func (s *StoreService) DeleteCategory(ctx context.Context, id uuid.UUID) (*int64, error) {
	dbResult, err := s.storage.Queries.DeleteCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	affected, err := dbResult.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affected, nil
}
