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

type CreateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategotyID  uuid.UUID
}

func (s *StoreService) CreateProduct(ctx context.Context, data CreateProductPayload) (*int64, error) {
	product, err := entity.NewProduct(data.Name, data.Description, data.Price, data.CategotyID)
	if err != nil {
		return nil, err
	}

	dbResult, err := s.storage.Queries.CreateProduct(ctx, store_storage.CreateProductParams{
		ID:          product.ID,
		Name:        product.Name,
		Slug:        product.Slug,
		Description: sql.NullString{String: *product.Description},
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
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

func (s *StoreService) GetProductById(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	return postgres.CallTx(ctx, s.storage.Db.Client, func(tx *sql.Tx) (*entity.Product, error) {
		queries := s.storage.Queries.WithTx(tx)

		dbResult, err := queries.GetProductById(ctx, id)

		if err != nil {
			return nil, err
		}

		return &entity.Product{
			ID:          dbResult.ID,
			Name:        dbResult.Name,
			Slug:        dbResult.Slug,
			Price:       dbResult.Price,
			CategoryID:  dbResult.CategoryID,
			Description: &dbResult.Description.String,
			CreatedAt:   dbResult.CreatedAt,
			UpdatedAt:   &dbResult.UpdatedAt.Int64,
		}, nil
	})
}

type GetManyProductParams struct {
	Page           int32
	PerPage        int32
	OrderBy        string
	OrderDirection string
}

func (s *StoreService) GetManyProduct(ctx context.Context, params GetManyProductParams) (*[]entity.Product, error) {
	return postgres.CallTx(ctx, s.storage.Db.Client, func(tx *sql.Tx) (*[]entity.Product, error) {
		queries := s.storage.Queries.WithTx(tx)

		pagination := postgres.Pagination(params.Page, params.PerPage)
		sorting := postgres.Sorting(params.OrderBy, params.OrderDirection)

		dbResult, err := queries.GetManyProduct(ctx, store_storage.GetManyProductParams{
			Limit:   pagination.Limit,
			Offset:  pagination.Offset,
			OrderBy: sorting,
		})

		if err != nil {
			return nil, err
		}

		var result []entity.Product = []entity.Product{}

		for _, dbProduct := range dbResult {
			result = append(result, entity.Product{
				ID:          dbProduct.ID,
				Name:        dbProduct.Name,
				Description: &dbProduct.Description.String,
				Price:       dbProduct.Price,
				CategoryID:  dbProduct.CategoryID,
				CreatedAt:   dbProduct.CreatedAt,
				UpdatedAt:   &dbProduct.UpdatedAt.Int64,
			})
		}

		return &result, nil
	})
}

type UpdateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uuid.UUID
}

func (s *StoreService) UpdateProduct(ctx context.Context, id uuid.UUID, data UpdateProductPayload) (*int64, error) {
	return postgres.CallTx(ctx, s.storage.Db.Client, func(tx *sql.Tx) (*int64, error) {
		queries := s.storage.Queries.WithTx(tx)

		res, err := queries.GetProductById(ctx, id)
		if err != nil {
			return nil, err
		}

		product := storage.ToProduct(&res)

		product.Update(data.Name, data.Description, data.Price, data.CategoryID)

		dbResult, err := queries.UpdateProduct(ctx, store_storage.UpdateProductParams{
			ID:          product.ID,
			Name:        product.Name,
			Slug:        product.Slug,
			Price:       product.Price,
			CategoryID:  product.CategoryID,
			Description: sql.NullString{String: data.Description},
			UpdatedAt:   sql.NullInt64{Int64: *product.UpdatedAt, Valid: true},
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

func (s *StoreService) DeleteProduct(ctx context.Context, id uuid.UUID) (*int64, error) {
	dbResult, err := s.storage.Queries.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	affected, err := dbResult.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affected, nil
}
