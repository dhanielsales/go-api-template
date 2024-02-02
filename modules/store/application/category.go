package application

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/dhanielsales/golang-scaffold/entity"
	appError "github.com/dhanielsales/golang-scaffold/internal/error"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"

	"github.com/dhanielsales/golang-scaffold/modules/store/storage"
	store_storage "github.com/dhanielsales/golang-scaffold/modules/store/storage/postgres"
)

type CreateCategoryPayload struct {
	Name        string
	Description string
	ImageUrl    string
}

func (s *StoreService) CreateCategory(ctx context.Context, data CreateCategoryPayload) (*int64, error) {
	return postgres.CallTx(ctx, s.storage.Postgres.Client, func(tx *sql.Tx) (*int64, error) {
		queries := s.storage.Queries.WithTx(tx)

		category, err := entity.NewCategory(data.Name, data.Description)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		payload := store_storage.CreateCategoryParams{
			ID:        category.ID,
			Name:      category.Name,
			Slug:      category.Slug,
			CreatedAt: category.CreatedAt,
		}

		if data.Description != "" {
			payload.Description = sql.NullString{String: data.Description, Valid: true}
		}

		dbResult, err := queries.CreateCategory(ctx, payload)

		if err != nil {
			if postgres.IsUniqueViolationByField(err, "slug") {
				return nil, appError.New(err, appError.UnprocessableEntityError, "Já existe uma categoria com esse slug")
			}

			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		affected, err := dbResult.RowsAffected()
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		err = s.storage.Cache.DeleteAllCategoryInCache(ctx)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		if _, err = s.external.Example.CreateImage(ctx, category.ID.String(), data.ImageUrl); err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		return &affected, nil
	})

}

func (s *StoreService) GetCategoryById(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	return postgres.CallTx(ctx, s.storage.Postgres.Client, func(tx *sql.Tx) (*entity.Category, error) {
		categoryInCache := s.storage.Cache.GetCategoryInCache(ctx, id)

		if categoryInCache != nil {
			return categoryInCache, nil
		}

		queries := s.storage.Queries.WithTx(tx)

		dbResult, err := queries.GetCategoryById(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, appError.New(err, appError.NotFoundError, "Category not found")
			}

			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		dbResultProd, err := queries.GetManyProductByCategoryId(ctx, id)
		if err != nil && err != sql.ErrNoRows {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
		}

		var products []entity.Product = []entity.Product{}

		for _, prod := range dbResultProd {
			curr := storage.ToProduct(&prod)
			products = append(products, *curr)
		}

		res := storage.ToCategory(&dbResult)
		res.Products = &products

		err = s.storage.Cache.SetCategoryInCache(ctx, *res, time.Hour*24)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

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
	return postgres.CallTx(ctx, s.storage.Postgres.Client, func(tx *sql.Tx) (*[]entity.Category, error) {
		queries := s.storage.Queries.WithTx(tx)

		pagination := postgres.Pagination(params.Page, params.PerPage)
		sorting := postgres.Sorting(params.OrderBy, params.OrderDirection)

		dbResult, err := queries.GetManyCategory(ctx, store_storage.GetManyCategoryParams{
			Limit:   pagination.Limit,
			Offset:  pagination.Offset,
			OrderBy: sorting,
		})

		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		var result []entity.Category = []entity.Category{}

		for _, dbCategory := range dbResult {
			curr := storage.ToCategory(&dbCategory)

			ext, err := s.external.Example.GetImage(ctx, curr.ID.String())
			if err != nil {
				return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
			}

			if len(ext.AllImages) > 0 {
				curr.ImageUrl = &ext.AllImages[0].Url
			}

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
	return postgres.CallTx(ctx, s.storage.Postgres.Client, func(tx *sql.Tx) (*int64, error) {
		queries := s.storage.Queries.WithTx(tx)

		res, err := queries.GetCategoryById(ctx, id)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		category := storage.ToCategory(&res)

		category.Update(data.Name, data.Description)

		payload := store_storage.UpdateCategoryParams{
			ID:        category.ID,
			Name:      category.Name,
			Slug:      category.Slug,
			UpdatedAt: sql.NullInt64{Int64: *category.UpdatedAt, Valid: true},
		}

		if data.Description != "" {
			payload.Description = sql.NullString{String: data.Description, Valid: true}
		}

		dbResult, err := queries.UpdateCategory(ctx, payload)
		if err != nil {
			if postgres.IsUniqueViolationByField(err, "slug") {
				return nil, appError.New(err, appError.UnprocessableEntityError, "Já existe uma categoria com esse slug")
			}

			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		affected, err := dbResult.RowsAffected()
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		err = s.storage.Cache.DeleteCategoryInCache(ctx, id)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		return &affected, nil
	})
}

func (s *StoreService) DeleteCategory(ctx context.Context, id uuid.UUID) (*int64, error) {
	dbResult, err := s.storage.Queries.DeleteCategory(ctx, id)
	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
	}

	affected, err := dbResult.RowsAffected()
	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
	}

	err = s.storage.Cache.DeleteCategoryInCache(ctx, id)
	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
	}

	return &affected, nil
}
