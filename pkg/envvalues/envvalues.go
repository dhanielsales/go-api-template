package envvalues

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/dhanielsales/go-api-template/pkg/logger"
)

const (
	ENV_TAG             = "env"
	DEFAULT_VALUE_TAG   = "default"
	DELIMITER_VALUE_TAG = "delimiter"
)

// Load returns a copy of the Values in the type parameter, filled by environment vars.
// The type parameter needs to be a struct type and have the proper annotations to identify
// environment vars.
//
// Annotation tags:
//
// Environment var name = "env"
//
// Default value = "default"
//
// Delimiter for string slices = "delimiter" (default delimiter ";")
//
// Example:
//
//	type MyEnvs struct {
//		String   string   `env:"string" default:"string"`
//		Strings  []string `env:"strings" default:"1;2;3"`
//		Strings2 []string `env:"strings" default:"1 2 3" delimiter:" "`
//		Bool     bool     `env:"bool" default:"true"`
//		Int      int      `env:"int" default:"12"`
//	}
func Load[Values any]() *Values {
	values := new(Values)

	config := reflect.Indirect(reflect.ValueOf(values))
	for i := range config.NumField() {
		field := config.Type().Field(i)
		envVar := field.Tag.Get(ENV_TAG)

		if envVar == "" {
			logger.Error("'env' tag not found in struct")
			continue
		}

		fieldValue := config.Field(i)
		valueOnEnv := os.Getenv(envVar)
		if valueOnEnv == "" {
			valueOnEnv = field.Tag.Get(DEFAULT_VALUE_TAG)
			logger.Debug("Env var not found, using default value", logger.LogString("envVar", envVar), logger.LogString("default", valueOnEnv))
		}

		switch fieldValue.Kind() {
		case reflect.Bool:
			value, err := strconv.ParseBool(valueOnEnv)
			if err != nil {
				logger.Error("Error on parse bool env var", logger.LogString("envVar", envVar), logger.LogErr("err", err))
				continue
			}
			fieldValue.SetBool(value)
		case reflect.Int:
			value, err := strconv.Atoi(valueOnEnv)
			if err != nil {
				logger.Error("Error on parse int env var", logger.LogString("envVar", envVar), logger.LogErr("err", err))
				continue
			}
			fieldValue.SetInt(int64(value))
		case reflect.String:
			fieldValue.SetString(valueOnEnv)
		case reflect.Slice:
			delimiter := field.Tag.Get(DELIMITER_VALUE_TAG)
			if delimiter == "" {
				delimiter = ";"
			}

			sliceValue := strings.Split(valueOnEnv, delimiter)
			slice := reflect.MakeSlice(fieldValue.Type(), len(sliceValue), len(sliceValue))
			for i, val := range sliceValue {
				slice.Index(i).Set(reflect.ValueOf(val))
			}

			fieldValue.Set(slice)
		default:
			logger.Error("Type not supported", logger.LogAny("kind", fieldValue.Kind()))
		}
	}

	return values
}
