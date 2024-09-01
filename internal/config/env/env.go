package env

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
)

type EnvVars struct {
	ENV               string `env:"ENV"`
	APP_NAME          string `env:"APP_NAME"`
	HTTP_ADDRESS      string `env:"HTTP_ADDRESS"`
	HTTP_PORT         string `env:"HTTP_PORT"`
	HTTP_ALLOW_ORIGIN string `env:"HTTP_ALLOW_ORIGIN"`
	POSTGRES_URL      string `env:"POSTGRES_URL"`
	REDIS_URL         string `env:"REDIS_URL"`

	KAFKA_BROKER                   string `env:"KAFKA_BROKER"`
	KAFKA_SSL                      bool   `env:"KAFKA_SSL"`
	KAFKA_USERNAME                 string `env:"KAFKA_USERNAME"`
	KAFKA_PASSWORD                 string `env:"KAFKA_PASSWORD"`
	KAFKA_GROUP_ID                 string `env:"KAFKA_GROUP_ID"`
	KAFKA_CLIENT_ID                string `env:"KAFKA_CLIENT_ID"`
	KAFKA_NOTIFICATION_ERROR_EMAIL string `env:"KAFKA_NOTIFICATION_ERROR_EMAIL"`
}

var once sync.Once
var instance *EnvVars

// GetInstance returns the singleton instance of EnvVars or panic if it's not load correctly
func GetInstance() *EnvVars {
	if instance == nil {
		once.Do(func() {
			instance = load()
		})
	}

	return instance
}

func load() *EnvVars {
	var configStruct EnvVars

	config := reflect.Indirect(reflect.ValueOf(&configStruct))
	for i := 0; i < config.NumField(); i++ {
		envVar := config.Type().Field(i).Tag.Get("env")

		if envVar == "" {
			log.Panic("'env' tag not found in struct")
		}

		field := config.Field(i)
		valueOnEnv := os.Getenv(envVar)
		if valueOnEnv == "" {
			log.Panicf("Env var '%s' not found", envVar)
		}

		switch field.Kind() {
		case reflect.Bool:
			value, err := strconv.ParseBool(valueOnEnv)
			if err != nil {
				log.Panicf("Error on parse bool env var '%s': %s", envVar, err.Error())
			}
			config.Field(i).SetBool(value)
		case reflect.Int:
			value, err := strconv.Atoi(valueOnEnv)
			if err != nil {
				log.Panicf("Error on parse int env var '%s': %s", envVar, err.Error())
			}
			config.Field(i).SetInt(int64(value))
		case reflect.String:
			config.Field(i).SetString(valueOnEnv)
		default:
			log.Panicf("Type '%s' not supported", field.Kind())
		}
	}

	return &configStruct
}
