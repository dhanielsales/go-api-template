package envvalues_test

import (
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/envvalues"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	t.Run("[ERROR] 'env' tag not found in struct", func(tt *testing.T) {
		tt.Parallel()

		type values struct {
			Name string
		}

		expected := &values{}
		result := envvalues.Load[values]()

		assert.Equal(tt, expected, result)
	})

	t.Run("[DEBUG] Env var not found, using default value", func(tt *testing.T) {
		tt.Parallel()

		type values struct {
			Name string `env:"name" default:"foo"`
		}

		expected := &values{
			Name: "foo",
		}
		result := envvalues.Load[values]()

		assert.Equal(tt, expected, result)
	})

	t.Run("[ERROR] Error on parse bool env var", func(tt *testing.T) {
		tt.Parallel()

		type values struct {
			Bool bool `env:"bool" default:"foo"`
		}

		expected := &values{}
		result := envvalues.Load[values]()

		assert.Equal(tt, expected, result)
	})

	t.Run("[ERROR] Error on parse int env var", func(tt *testing.T) {
		tt.Parallel()

		type values struct {
			Int int `env:"int" default:"foo"`
		}

		expected := &values{}
		result := envvalues.Load[values]()

		assert.Equal(tt, expected, result)
	})

	t.Run("[ERROR] Type not supported", func(tt *testing.T) {
		tt.Parallel()

		type values struct {
			Map map[string]any `env:"map" default:"{}"`
		}

		expected := &values{}
		result := envvalues.Load[values]()

		assert.Equal(tt, expected, result)
	})

	t.Run("Success", func(tt *testing.T) {
		tt.Parallel()

		type values struct {
			String string `env:"string" default:"string"`
			Bool   bool   `env:"bool" default:"true"`
			Int    int    `env:"int" default:"12"`
		}

		expected := &values{
			String: "string",
			Bool:   true,
			Int:    12,
		}
		result := envvalues.Load[values]()

		assert.Equal(tt, expected, result)
	})
}
