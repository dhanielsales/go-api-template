package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages/category"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
	"github.com/dhanielsales/go-api-template/pkg/testutils"
	"github.com/dhanielsales/go-api-template/pkg/utils"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository *category.CategoryRepository
	}

	type args struct {
		data *models.Category
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
			name: "Error creating category - error CreateCategory method",
			fields: &fields{
				repository: newCategoryRepository(t, func(mocks *mocks) {
					mocks.storage.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(nil, errors.New("error creating category"))
				}),
			},
			args: &args{
				data: &models.Category{
					Name:        "Name",
					Slug:        "name",
					Description: utils.ToPtr("Description"),
				},
			},
			expected: &expected{
				err: errors.New("error creating category"),
			},
		},
		{
			name: "Error creating category - error RowsAffected method",
			fields: &fields{
				repository: newCategoryRepository(t, func(mocks *mocks) {
					result := sqlutils.NewMockResult(mocks.ctrl)
					result.EXPECT().RowsAffected().Return(int64(0), errors.New("error rows affected"))
					mocks.storage.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(result, nil)
				}),
			},
			args: &args{
				data: &models.Category{
					Name:        "Name",
					Slug:        "name",
					Description: utils.ToPtr("Description"),
				},
			},
			expected: &expected{
				err: errors.New("error rows affected"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				repository: newCategoryRepository(t, func(mocks *mocks) {
					result := sqlutils.NewMockResult(mocks.ctrl)
					result.EXPECT().RowsAffected().Return(int64(1), nil)
					mocks.storage.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(result, nil)
				}),
			},
			args: &args{
				data: &models.Category{
					Name:        "Name",
					Slug:        "name",
					Description: utils.ToPtr("Description"),
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
			affects, err := tt.fields.repository.CreateCategory(ctx, tt.args.data)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
