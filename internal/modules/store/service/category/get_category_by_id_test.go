package category_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCategoryByID(t *testing.T) {
	t.Parallel()

	type fields struct {
		service category.CategoryService
	}

	type args struct {
		id uuid.UUID
	}

	type expected struct {
		data *models.Category
		err  error
	}

	ctx := context.Background()
	id := uuid.New()

	tests := []struct {
		name     string
		fields   *fields
		args     *args
		expected *expected
	}{
		{
			name: "Success - fiding in cache",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryInCache(gomock.Any(), gomock.Any()).Return(&models.Category{
						ID:   id,
						Name: "Name",
						Slug: "name",
					})
				}),
			},
			args: &args{id},
			expected: &expected{
				err: nil,
				data: &models.Category{
					ID:   id,
					Name: "Name",
					Slug: "name",
				},
			},
		},
		{
			name: "Error getting category - general error",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryInCache(gomock.Any(), gomock.Any()).Return(nil)
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error deleting category"))
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(errors.New("error deleting category")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Error getting category - not found",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryInCache(gomock.Any(), gomock.Any()).Return(nil)
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(sql.ErrNoRows).WithDescription("category not found").WithStatusCode(http.StatusNotFound),
			},
		},
		{
			name: "Error setting category in cache",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryInCache(gomock.Any(), gomock.Any()).Return(nil)
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(&models.Category{
						ID:   id,
						Name: "Name",
						Slug: "name",
					}, nil)
					mocks.categoryRepository.EXPECT().SetCategoryInCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error setting category in cache"))
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(errors.New("error setting category in cache")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryInCache(gomock.Any(), gomock.Any()).Return(nil)
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(&models.Category{
						ID:   id,
						Name: "Name",
						Slug: "name",
					}, nil)
					mocks.categoryRepository.EXPECT().SetCategoryInCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				}),
			},
			args: &args{id},
			expected: &expected{
				err: nil,
				data: &models.Category{
					ID:   id,
					Name: "Name",
					Slug: "name",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.GetCategoryByID(ctx, tt.args.id)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
