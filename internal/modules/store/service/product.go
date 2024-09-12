package service

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/postgres"

	"github.com/dhanielsales/go-api-template/internal/models"
)

type CreateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategotyID  uuid.UUID
}

func (s *StoreService) CreateProduct(ctx context.Context, data CreateProductPayload) (int64, error) {
	product, err := models.NewProduct(data.Name, data.Description, data.Price, data.CategotyID)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("Can't processable product entity").WithStatusCode(http.StatusUnprocessableEntity)
	}

	affected, err := s.repository.Persistence.CreateProduct(ctx, product)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("Can't processable product entity").WithStatusCode(http.StatusUnprocessableEntity)
	}

	return affected, nil
}

func (s *StoreService) GetProductByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*models.Product, error) {
		queries := s.repository.Persistence.WithTx(tx)

		product, err := queries.GetProductByID(ctx, id)
		if err != nil {
			return nil, apperror.FromError(err).WithDescription("Can't processable product entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		return product, nil
	})
}

type GetManyProductParams struct {
	Page           string
	PerPage        string
	OrderBy        string
	OrderDirection string
}

func (s *StoreService) GetManyProduct(ctx context.Context, params GetManyProductParams) ([]*models.Product, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) ([]*models.Product, error) {
		queries := s.repository.Persistence.WithTx(tx)

		products, err := queries.GetManyProduct(ctx, models.GetManyProductPayload{
			Page:           params.Page,
			PerPage:        params.PerPage,
			OrderBy:        params.OrderBy,
			OrderDirection: params.OrderDirection,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, apperror.FromError(err).WithDescription("Product not found").WithStatusCode(http.StatusNotFound)
			}

			return nil, apperror.FromError(err).WithDescription("Can't processable product entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		return products, nil
	})
}

type UpdateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uuid.UUID
}

func (s *StoreService) UpdateProduct(ctx context.Context, id uuid.UUID, data UpdateProductPayload) (int64, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (int64, error) {
		queries := s.repository.Persistence.WithTx(tx)

		product, err := queries.GetProductByID(ctx, id)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("Can't processable product entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		product.Update(data.Name, data.Description, data.Price, data.CategoryID)

		affected, err := queries.UpdateProduct(ctx, id, product)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("Can't processable product entity").WithStatusCode(http.StatusUnprocessableEntity)
		}

		return affected, nil
	})
}

func (s *StoreService) DeleteProduct(ctx context.Context, id uuid.UUID) (int64, error) {
	affected, err := s.repository.Persistence.DeleteProduct(ctx, id)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("Can't processable product entity").WithStatusCode(http.StatusUnprocessableEntity)
	}

	return affected, nil
}
