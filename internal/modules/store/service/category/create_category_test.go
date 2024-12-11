package category_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		service category.CategoryService
	}

	type args struct {
		data category.CreateCategoryPayload
	}

	type expected struct {
		data int64
		err  error
	}

	ctx := context.Background()

	tests := []struct {
		name     string
		fields   *fields
		args     *args
		expected *expected
	}{
		{
			name: "Error invalid payload",
			fields: &fields{
				service: newCategoryService(t, nil),
			},
			args: &args{
				data: category.CreateCategoryPayload{
					Name:        "",
					Description: "",
				},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("name is required")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Error creating category - unique violation",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(0), &pq.Error{Message: "duplicate key value violates unique constraint 'slug'"})
				}),
			},
			args: &args{
				data: category.CreateCategoryPayload{
					Name:        "Name",
					Description: "Description",
				},
			},
			expected: &expected{
				err: apperror.FromError(&pq.Error{Message: "duplicate key value violates unique constraint 'slug'"}).WithDescription("category with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity),
			},
		},
		{
			name: "Error creating category - general error",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error creating category"))
				}),
			},
			args: &args{
				data: category.CreateCategoryPayload{
					Name:        "Name",
					Description: "Description",
				},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error creating category")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Error cleaning category cache",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(1), nil)
					mocks.repository.EXPECT().DeleteAllCategoriesInCache(gomock.Any()).Return(errors.New("error cleaning cache"))
				}),
			},
			args: &args{
				data: category.CreateCategoryPayload{
					Name:        "Name",
					Description: "Description",
				},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error cleaning cache")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(1), nil)
					mocks.repository.EXPECT().DeleteAllCategoriesInCache(gomock.Any()).Return(nil)
				}),
			},
			args: &args{
				data: category.CreateCategoryPayload{
					Name:        "Name",
					Description: "Description",
				},
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
			affects, err := tt.fields.service.CreateCategory(ctx, tt.args.data)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
