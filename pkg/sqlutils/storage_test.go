package sqlutils_test

import (
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStorageBootstrapSuccess(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mocksql := sqlutils.NewMockSQLDB(ctrl)

	s := sqlutils.New(mocksql)
	assert.NotNil(t, s.Client)
}

func TestStorageCleanupSuccess(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mocksql := sqlutils.NewMockSQLDB(ctrl)

	s := sqlutils.New(mocksql)
	mocksql.EXPECT().Close().Return(nil)

	assert.NoError(t, s.Cleanup())
}

func TestStorageCleanupError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mocksql := sqlutils.NewMockSQLDB(ctrl)

	s := sqlutils.New(mocksql)
	mocksql.EXPECT().Close().Return(errors.New("Error on close"))
	err := s.Cleanup()

	assert.Error(t, err)
	assert.EqualError(t, err, "error closing postgress connection: Error on close")
}
