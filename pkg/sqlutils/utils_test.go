package sqlutils_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
	"github.com/dhanielsales/go-api-template/pkg/testutils"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestPagination(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		page     string
		perPage  string
		expected sqlutils.PaginationResult
	}{
		{"ValidInput", "2", "20", sqlutils.PaginationResult{Limit: 20, Offset: 20}},
		{"InvalidPage", "invalid", "20", sqlutils.PaginationResult{Limit: 20, Offset: 0}},
		{"InvalidPerPage", "2", "invalid", sqlutils.PaginationResult{Limit: 10, Offset: 10}},
		{"ZeroPage", "0", "20", sqlutils.PaginationResult{Limit: 20, Offset: 0}},
		{"ZeroPerPage", "2", "0", sqlutils.PaginationResult{Limit: 10, Offset: 10}},
		{"LargePerPage", "2", "150", sqlutils.PaginationResult{Limit: 100, Offset: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := sqlutils.Pagination(tt.page, tt.perPage)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSorting(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		field     string
		direction string
		expected  string
	}{
		{"ValidInput", "name", "ASC", "name ASC"},
		{"EmptyField", "", "DESC", "created_at DESC"},
		{"EmptyDirection", "name", "", "name DESC"},
		{"InvalidDirection", "name", "invalid", "name DESC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := sqlutils.Sorting(tt.field, tt.direction)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWithTx(t *testing.T) {
	t.Parallel()

	type expected struct {
		data any
		err  error
	}

	tests := []struct {
		name      string
		txHandler func(tx sqlutils.SQLTX) (any, error)
		prepare   func() sqlutils.SQLDB
		expected  *expected
	}{
		{
			name:      "Success",
			txHandler: func(tx sqlutils.SQLTX) (any, error) { return 1, nil },
			prepare: func() sqlutils.SQLDB {
				ctrl := gomock.NewController(t)
				mockdb := sqlutils.NewMockSQLDB(ctrl)
				mocktx := sqlutils.NewMockSQLTX(ctrl)

				mockdb.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(mocktx, nil)
				mocktx.EXPECT().Commit().Return(nil)

				return mockdb
			},
			expected: &expected{
				data: 1,
				err:  nil,
			},
		},
		{
			name:      "Success rollback",
			txHandler: func(tx sqlutils.SQLTX) (any, error) { return 0, errors.New("error in handler") },
			prepare: func() sqlutils.SQLDB {
				ctrl := gomock.NewController(t)
				mockdb := sqlutils.NewMockSQLDB(ctrl)
				mocktx := sqlutils.NewMockSQLTX(ctrl)

				mockdb.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(mocktx, nil)
				mocktx.EXPECT().Rollback().Return(nil)

				return mockdb
			},
			expected: &expected{
				data: nil,
				err:  errors.New("error in handler"),
			},
		},
		{
			name:      "Success commit",
			txHandler: func(tx sqlutils.SQLTX) (any, error) { return 1, nil },
			prepare: func() sqlutils.SQLDB {
				ctrl := gomock.NewController(t)
				mockdb := sqlutils.NewMockSQLDB(ctrl)
				mocktx := sqlutils.NewMockSQLTX(ctrl)

				mockdb.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(mocktx, nil)
				mocktx.EXPECT().Commit().Return(nil)

				return mockdb
			},
			expected: &expected{
				data: 1,
				err:  nil,
			},
		},
		{
			name:      "Error in the BeginTx",
			txHandler: func(tx sqlutils.SQLTX) (any, error) { return 1, nil },
			prepare: func() sqlutils.SQLDB {
				ctrl := gomock.NewController(t)
				mockdb := sqlutils.NewMockSQLDB(ctrl)

				mockdb.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(nil, errors.New("error begin tx"))

				return mockdb
			},
			expected: &expected{
				data: nil,
				err:  errors.New("error begin tx"),
			},
		},
		{
			name:      "Error in the commit",
			txHandler: func(tx sqlutils.SQLTX) (any, error) { return 1, nil },
			prepare: func() sqlutils.SQLDB {
				ctrl := gomock.NewController(t)
				mockdb := sqlutils.NewMockSQLDB(ctrl)
				mocktx := sqlutils.NewMockSQLTX(ctrl)

				mockdb.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(mocktx, nil)
				mocktx.EXPECT().Commit().Return(errors.New("error commit"))

				return mockdb
			},
			expected: &expected{
				data: nil,
				err:  errors.New("error commit"),
			},
		},
		{
			name:      "Error in the rollback",
			txHandler: func(tx sqlutils.SQLTX) (any, error) { return 0, errors.New("error in the handler") },
			prepare: func() sqlutils.SQLDB {
				ctrl := gomock.NewController(t)
				mockdb := sqlutils.NewMockSQLDB(ctrl)
				mocktx := sqlutils.NewMockSQLTX(ctrl)

				mockdb.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(mocktx, nil)
				mocktx.EXPECT().Rollback().Return(errors.New("error in the rollback"))

				return mockdb
			},
			expected: &expected{
				data: nil,
				err:  errors.Join(errors.New("error in the handler"), errors.New("error in the rollback")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db := tt.prepare()
			result, err := sqlutils.WithTx(context.Background(), db, tt.txHandler)

			assert.Equal(t, tt.expected.data, result)
			testutils.ErrorEqual(t, tt.expected.err, err)
		})
	}
}
