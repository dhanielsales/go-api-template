package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetManyCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		service category.CategoryService
	}

	type args struct {
		params category.GetManyCategoryParams
	}

	type expected struct {
		data []*models.Category
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
			name: "Error getting categories - general error",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetManyCategory(gomock.Any(), gomock.Any()).Return(nil, errors.New("error getting categories"))
				}),
			},
			args: &args{
				params: category.GetManyCategoryParams{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error getting categories")).WithDescription("can't process category entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newCategoryService(t, func(mocks *mocks) {
					mocks.categoryRepository.EXPECT().GetManyCategory(gomock.Any(), gomock.Any()).Return([]*models.Category{{
						ID:   id,
						Name: "Name",
						Slug: "name",
					}}, nil)
				}),
			},
			args: &args{
				params: category.GetManyCategoryParams{},
			},
			expected: &expected{
				err: nil,
				data: []*models.Category{{
					ID:   id,
					Name: "Name",
					Slug: "name",
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.GetManyCategory(ctx, tt.args.params)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
