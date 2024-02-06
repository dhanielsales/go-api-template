package application

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/dhanielsales/golang-scaffold/entity"
	appError "github.com/dhanielsales/golang-scaffold/internal/error"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
)

type CreateCategoryPayload struct {
	Name        string
	Description string
	ImageUrl    string
}

func (s *StoreService) CreateCategory(ctx context.Context, data CreateCategoryPayload) (*int64, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*int64, error) {
		queries := s.repository.Persistence.WithTx(tx)

		category, err := entity.NewCategory(data.Name, data.Description)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		affecteds, err := queries.CreateCategory(ctx, category)

		if err != nil {
			if postgres.IsUniqueViolationByField(err, "slug") {
				return nil, appError.New(err, appError.UnprocessableEntityError, "Já existe uma categoria com esse slug")
			}

			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		err = s.repository.Cache.DeleteAllCategoryInCache(ctx)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		if _, err = s.external.Example.CreateImage(ctx, category.ID.String(), data.ImageUrl); err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		return affecteds, nil
	})

}

func (s *StoreService) GetCategoryById(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*entity.Category, error) {
		categoryInCache := s.repository.Cache.GetCategoryInCache(ctx, id)

		if categoryInCache != nil {
			return categoryInCache, nil
		}

		queries := s.repository.Persistence.WithTx(tx)

		category, err := queries.GetCategoryById(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, appError.New(err, appError.NotFoundError, "Category not found")
			}

			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		err = s.repository.Cache.SetCategoryInCache(ctx, *category, time.Hour*24)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		return category, nil
	})
}

type GetManyCategoryParams struct {
	Page           string
	PerPage        string
	OrderBy        string
	OrderDirection string
}

func (s *StoreService) GetManyCategory(ctx context.Context, params GetManyCategoryParams) (*[]entity.Category, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*[]entity.Category, error) {
		queries := s.repository.Persistence.WithTx(tx)

		dbResult, err := queries.GetManyCategory(ctx, entity.GetManyCategoryPayload{
			Page:           params.Page,
			PerPage:        params.PerPage,
			OrderBy:        params.OrderBy,
			OrderDirection: params.OrderDirection,
		})

		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		var result []entity.Category = []entity.Category{}

		for _, dbCategory := range *dbResult {
			ext, err := s.external.Example.GetImage(ctx, dbCategory.ID.String())
			if err != nil {
				return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
			}

			if len(ext.AllImages) > 0 {
				dbCategory.ImageUrl = &ext.AllImages[0].Url
			}

			result = append(result, dbCategory)
		}

		return &result, nil
	})
}

type UpdateCategoryPayload struct {
	Name        string
	Description string
}

func (s *StoreService) UpdateCategory(ctx context.Context, id uuid.UUID, data UpdateCategoryPayload) (*int64, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*int64, error) {
		queries := s.repository.Persistence.WithTx(tx)

		category, err := queries.GetCategoryById(ctx, id)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		category.Update(data.Name, data.Description)

		affected, err := queries.UpdateCategory(ctx, id, category)
		if err != nil {
			if postgres.IsUniqueViolationByField(err, "slug") {
				return nil, appError.New(err, appError.UnprocessableEntityError, "Já existe uma categoria com esse slug")
			}

			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		err = s.repository.Cache.DeleteCategoryInCache(ctx, id)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
		}

		return affected, nil
	})
}

func (s *StoreService) DeleteCategory(ctx context.Context, id uuid.UUID) (*int64, error) {
	affected, err := s.repository.Persistence.DeleteCategory(ctx, id)
	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
	}

	err = s.repository.Cache.DeleteCategoryInCache(ctx, id)
	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable category entity")
	}

	return affected, nil
}

func (s *StoreService) GetManyCategoryNoDb(ctx context.Context) (*[]entity.Category, error) {
	return s.repository.Fake.GetManyCategory(ctx, entity.GetManyCategoryPayload{})
}
