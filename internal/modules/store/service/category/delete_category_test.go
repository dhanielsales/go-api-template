package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		service category.CategoryService
	}

	type args struct {
		id uuid.UUID
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
			name: "Error deleting category - general error",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().DeleteCategory(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error deleting category"))
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(errors.New("error deleting category")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Error deleting category cache",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().DeleteCategory(gomock.Any(), gomock.Any()).Return(int64(1), nil)
					mocks.categoryRepository.EXPECT().DeleteCategoryInCache(gomock.Any(), gomock.Any()).Return(errors.New("error cleaning cache"))
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(errors.New("error cleaning cache")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().DeleteCategory(gomock.Any(), gomock.Any()).Return(int64(1), nil)
					mocks.categoryRepository.EXPECT().DeleteCategoryInCache(gomock.Any(), gomock.Any()).Return(nil)
				}),
			},
			args: &args{id},
			expected: &expected{
				err:  nil,
				data: int64(1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.DeleteCategory(ctx, tt.args.id)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
