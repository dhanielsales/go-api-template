package apperror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	t.Parallel()

	t.Run("test (apperror).WithDescription method", func(tt *testing.T) {
		tt.Parallel()

		err := apperror.New("description")
		assert.Equal(tt, "description", err.Description)

		err.WithDescription("changed")
		assert.Equal(tt, "changed", err.Description)
	})

	t.Run("test (apperror).WithStatusCode method", func(tt *testing.T) {
		tt.Parallel()

		err := apperror.New("description")
		assert.Equal(tt, http.StatusText(http.StatusInternalServerError), err.Name)
		assert.Equal(tt, http.StatusInternalServerError, err.StatusCode())

		err.WithStatusCode(http.StatusBadGateway)
		assert.Equal(tt, http.StatusText(http.StatusBadGateway), err.Name)
		assert.Equal(tt, http.StatusBadGateway, err.StatusCode())
	})

	t.Run("test (apperror).WithStatusCode method invalid status", func(tt *testing.T) {
		tt.Parallel()

		err := apperror.New("description")
		assert.Equal(tt, http.StatusText(http.StatusInternalServerError), err.Name)
		assert.Equal(tt, http.StatusInternalServerError, err.StatusCode())

		err.WithStatusCode(0)
		assert.Equal(tt, "", err.Name)
		assert.Equal(tt, http.StatusInternalServerError, err.StatusCode())
	})

	t.Run("test (apperror).WithDetails method", func(tt *testing.T) {
		tt.Parallel()

		err := apperror.New("description")
		assert.Equal(tt, nil, err.Details)
		err.WithDetails("details")
		assert.Equal(tt, "details", err.Details)
	})

	t.Run("test (apperror).Merge method", func(tt *testing.T) {
		tt.Parallel()

		err := apperror.New("description")
		err.Merge(errors.New("foo"))
		assert.Equal(tt, "error Internal Server Error: err foo - description", err.Error())
	})

	t.Run("test (apperror).Unwrap method", func(tt *testing.T) {
		tt.Parallel()

		sourceErr := errors.New("foo")
		err := apperror.FromError(sourceErr)
		assert.EqualError(tt, sourceErr, err.Unwrap().Error())
	})

	t.Run("test (apperror).Stack method", func(tt *testing.T) {
		tt.Parallel()

		err := apperror.New("description")
		err.Merge(errors.New("foo"))
		assert.Equal(tt, "error Internal Server Error: err foo - description", err.Error())
		assert.NotEmpty(tt, err.Stack())
	})

	t.Run("test (apperror).Level method", func(tt *testing.T) {
		tt.Parallel()

		err := apperror.New("description")
		err.Merge(errors.New("foo"))
		assert.Equal(tt, err.Level.String(), apperror.Error.String())

		err.WithStatusCode(http.StatusBadRequest)
		assert.Equal(tt, err.Level.String(), apperror.Warn.String())

		err.WithStatusCode(http.StatusOK)
		assert.Equal(tt, err.Level.String(), apperror.Info.String())
	})
}
