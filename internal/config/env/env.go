package env

import (
	"sync"

	"github.com/dhanielsales/go-api-template/pkg/envvalues"
)

type Values struct {
	ENV               string `env:"ENV" default:"development"`
	APP_NAME          string `env:"APP_NAME" default:"go-api-template"`
	HTTP_ADDRESS      string `env:"HTTP_ADDRESS" default:"localhost"`
	HTTP_PORT         string `env:"HTTP_PORT" default:"8080"`
	HTTP_ALLOW_ORIGIN string `env:"HTTP_ALLOW_ORIGIN" default:"*"`
	POSTGRES_URL      string `env:"POSTGRES_URL" default:"postgres://postgres:postgres@localhost:5432/main?sslmode=disable"`
	REDIS_URL         string `env:"REDIS_URL" default:"redis://default:password@localhost:6379/0"`

	KAFKA_BROKER                   string `env:"KAFKA_BROKER" default:""`
	KAFKA_SSL                      bool   `env:"KAFKA_SSL" default:"false"`
	KAFKA_USERNAME                 string `env:"KAFKA_USERNAME" default:""`
	KAFKA_PASSWORD                 string `env:"KAFKA_PASSWORD" default:""`
	KAFKA_GROUP_ID                 string `env:"KAFKA_GROUP_ID" default:""`
	KAFKA_CLIENT_ID                string `env:"KAFKA_CLIENT_ID" default:""`
	KAFKA_NOTIFICATION_ERROR_EMAIL string `env:"KAFKA_NOTIFICATION_ERROR_EMAIL" default:"aaaaaa"`
}

var (
	once     sync.Once
	instance *Values
)

// GetInstance returns the singleton instance of EnvVars or panic if it's not load correctly
func GetInstance() *Values {
	if instance == nil {
		once.Do(func() {
			instance = envvalues.Load[Values]()
		})
	}

	return instance
}
