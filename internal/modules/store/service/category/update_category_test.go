package category_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"
	"github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		service category.CategoryService
	}

	type args struct {
		id     uuid.UUID
		params category.UpdateCategoryPayload
	}

	type expected struct {
		data int64
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
			name: "Error update category - get category to update",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error getting category"))
				}),
			},
			args: &args{
				params: category.UpdateCategoryPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error getting category")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Error update category - update category unique violation",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(&models.Category{ID: id}, nil)
					mocks.categoryRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), &pq.Error{Message: "duplicate key value violates unique constraint 'slug'"})
				}),
			},
			args: &args{
				params: category.UpdateCategoryPayload{},
			},
			expected: &expected{
				err: apperror.FromError(&pq.Error{Message: "duplicate key value violates unique constraint 'slug'"}).WithDescription("category with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity),
			},
		},
		{
			name: "Error update category - update category general error",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(&models.Category{ID: id}, nil)
					mocks.categoryRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error updating category"))
				}),
			},
			args: &args{
				params: category.UpdateCategoryPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error updating category")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Error update category - update category general error",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(&models.Category{ID: id}, nil)
					mocks.categoryRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error updating category"))
				}),
			},
			args: &args{
				params: category.UpdateCategoryPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error updating category")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Error update category - setting category in cache",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(&models.Category{ID: id}, nil)
					mocks.categoryRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil)
					mocks.categoryRepository.EXPECT().DeleteCategoryInCache(gomock.Any(), gomock.Any()).Return(errors.New("error deletting category in cache"))
				}),
			},
			args: &args{
				params: category.UpdateCategoryPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error deletting category in cache")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(&models.Category{ID: id}, nil)
					mocks.categoryRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil)
					mocks.categoryRepository.EXPECT().DeleteCategoryInCache(gomock.Any(), gomock.Any()).Return(nil)
				}),
			},
			args: &args{
				params: category.UpdateCategoryPayload{},
			},
			expected: &expected{
				err:  nil,
				data: int64(1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.UpdateCategory(ctx, tt.args.id, tt.args.params)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
