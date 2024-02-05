package application

import (
	"database/sql"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/dhanielsales/golang-scaffold/entity"
	appError "github.com/dhanielsales/golang-scaffold/internal/error"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
)

type CreateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategotyID  uuid.UUID
}

func (s *StoreService) CreateProduct(ctx context.Context, data CreateProductPayload) (*int64, error) {
	product, err := entity.NewProduct(data.Name, data.Description, data.Price, data.CategotyID)
	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
	}

	affected, err := s.repository.Persistence.CreateProduct(ctx, product)

	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
	}

	return affected, nil
}

func (s *StoreService) GetProductById(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*entity.Product, error) {
		queries := s.repository.Persistence.WithTx(tx)

		product, err := queries.GetProductById(ctx, id)

		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
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

func (s *StoreService) GetManyProduct(ctx context.Context, params GetManyProductParams) (*[]entity.Product, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*[]entity.Product, error) {
		queries := s.repository.Persistence.WithTx(tx)

		products, err := queries.GetManyProduct(ctx, entity.GetManyProductPayload{
			Page:           params.Page,
			PerPage:        params.PerPage,
			OrderBy:        params.OrderBy,
			OrderDirection: params.OrderDirection,
		})

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, appError.New(err, appError.NotFoundError, "Product not found")
			}

			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
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

func (s *StoreService) UpdateProduct(ctx context.Context, id uuid.UUID, data UpdateProductPayload) (*int64, error) {
	return postgres.CallTx(ctx, s.repository.Postgres.Client, func(tx *sql.Tx) (*int64, error) {
		queries := s.repository.Persistence.WithTx(tx)

		product, err := queries.GetProductById(ctx, id)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
		}

		product.Update(data.Name, data.Description, data.Price, data.CategoryID)

		affected, err := queries.UpdateProduct(ctx, id, product)
		if err != nil {
			return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
		}

		return affected, nil
	})
}

func (s *StoreService) DeleteProduct(ctx context.Context, id uuid.UUID) (*int64, error) {
	affected, err := s.repository.Persistence.DeleteProduct(ctx, id)
	if err != nil {
		return nil, appError.New(err, appError.UnprocessableEntityError, "Can't processable product entity")
	}

	return affected, nil
}
