package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/category"
	"github.com/dhanielsales/go-api-template/pkg/testutils"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func TestGetManyCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		controller *category.CategoryController
	}

	type args struct {
		c *testutils.EchoContext
	}

	type expected struct {
		err     error
		payload []byte
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
			name: "Error getting many categories",
			fields: &fields{
				controller: newCategoryController(t, func(mocks *mocks) {
					mocks.categoryService.EXPECT().GetManyCategory(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
				}),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, nil),
			},
			expected: &expected{
				err: errors.New("error"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				controller: newCategoryController(t, func(mocks *mocks) {
					mocks.categoryService.EXPECT().GetManyCategory(gomock.Any(), gomock.Any()).Return([]*models.Category{{
						ID:   id,
						Name: "Name",
						Slug: "name",
					}}, nil)
				}),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, nil),
			},
			expected: &expected{
				payload: testutils.ToByte([]*models.Category{{
					ID:   id,
					Name: "Name",
					Slug: "name",
				}}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.fields.controller.GetManyCategory(tt.args.c)

			testutils.ErrorEqual(t, tt.expected.err, err)
			testutils.BytesEqual(t, tt.expected.payload, tt.args.c.Rec.Body.Bytes())
		})
	}
}
