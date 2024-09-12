package service

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/postgres"

	"github.com/dhanielsales/go-api-template/internal/models"

	"github.com/google/uuid"
)

type CreateCategoryPayload struct {
	Name        string
	Description string
}

func (s *StoreService) CreateCategory(ctx context.Context, data CreateCategoryPayload) (int64, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (int64, error) {
		queries := s.repository.Persistence.WithTx(tx)

		category, err := models.NewCategory(data.Name, data.Description)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		affecteds, err := queries.CreateCategory(ctx, category)
		if err != nil {
			if postgres.IsUniqueViolationByField(err, "slug") {
				return 0, apperror.FromError(err).WithDescription("Category with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity)
			}

			return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		err = s.repository.Cache.DeleteAllCategoryInCache(ctx)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		return affecteds, nil
	})
}

func (s *StoreService) GetCategoryByID(ctx context.Context, id uuid.UUID) (*models.Category, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*models.Category, error) {
		categoryInCache := s.repository.Cache.GetCategoryInCache(ctx, id)

		if categoryInCache != nil {
			return categoryInCache, nil
		}

		queries := s.repository.Persistence.WithTx(tx)

		category, err := queries.GetCategoryByID(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, apperror.FromError(err).WithDescription("Category not found").WithStatusCode(http.StatusNotFound)
			}

			return nil, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		err = s.repository.Cache.SetCategoryInCache(ctx, category, time.Hour*24)
		if err != nil {
			return nil, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
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

func (s *StoreService) GetManyCategory(ctx context.Context, params GetManyCategoryParams) ([]*models.Category, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) ([]*models.Category, error) {
		queries := s.repository.Persistence.WithTx(tx)

		dbResult, err := queries.GetManyCategory(ctx, models.GetManyCategoryPayload{
			Page:           params.Page,
			PerPage:        params.PerPage,
			OrderBy:        params.OrderBy,
			OrderDirection: params.OrderDirection,
		})
		if err != nil {
			return nil, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		return dbResult, nil
	})
}

type UpdateCategoryPayload struct {
	Name        string
	Description string
}

func (s *StoreService) UpdateCategory(ctx context.Context, id uuid.UUID, data UpdateCategoryPayload) (int64, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (int64, error) {
		queries := s.repository.Persistence.WithTx(tx)

		category, err := queries.GetCategoryByID(ctx, id)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		category.Update(data.Name, data.Description)

		affected, err := queries.UpdateCategory(ctx, id, category)
		if err != nil {
			if postgres.IsUniqueViolationByField(err, "slug") {
				return 0, apperror.FromError(err).WithDescription("Category with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity)
			}

			return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		err = s.repository.Cache.DeleteCategoryInCache(ctx, id)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		return affected, nil
	})
}

func (s *StoreService) DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error) {
	affected, err := s.repository.Persistence.DeleteCategory(ctx, id)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
	}

	err = s.repository.Cache.DeleteCategoryInCache(ctx, id)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("Can't processable category entity").WithStatusCode(http.StatusUnprocessableEntity)
	}

	return affected, nil
}
