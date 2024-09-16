package envvalues

import (
	"os"
	"reflect"
	"strconv"

	"github.com/dhanielsales/go-api-template/pkg/logger"
)

const (
	ENV_TAG           = "env"
	DEFAULT_VALUE_TAG = "default"
)

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
			config.Field(i).SetBool(value)
		case reflect.Int:
			value, err := strconv.Atoi(valueOnEnv)
			if err != nil {
				logger.Error("Error on parse int env var", logger.LogString("envVar", envVar), logger.LogErr("err", err))
				continue
			}
			config.Field(i).SetInt(int64(value))
		case reflect.String:
			config.Field(i).SetString(valueOnEnv)
		default:
			logger.Error("Type not supported", logger.LogAny("kind", fieldValue.Kind()))
		}
	}

	return values
}
