package env

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type EnvVars struct {
	APP_NAME     string `mapstructure:"APP_NAME"`
	MONGODB_URL  string `mapstructure:"MONGODB_URL"`
	MONGODB_NAME string `mapstructure:"MONGODB_NAME"`
	POSTGRES_URL string `mapstructure:"POSTGRES_URL"`
	REDIS_URL    string `mapstructure:"REDIS_URL"`
	PORT         string `mapstructure:"PORT"`
	EXTERNAL_URL string `mapstructure:"EXTERNAL_URL"`
}

func LoadEnv() (*EnvVars, error) {
	fmt.Println("LoadEnv")
	var config EnvVars
	fmt.Println("config", config)

	env := os.Getenv("GO_ENV")
	fmt.Println("env", env)
	if env == "production" {
		fmt.Println("LoadEnv if")

		val := reflect.Indirect(reflect.ValueOf(&config))
		for i := 0; i < val.NumField(); i++ {
			envVar := val.Type().Field(i).Tag.Get("mapstructure")
			fmt.Println("envVar", envVar)

			if envVar == "" {
				return nil, fmt.Errorf("Env var '%s' not found", envVar)
			}

			fmt.Println("os.Getenv(envVar)", os.Getenv(envVar))

			val.Field(i).SetString(os.Getenv(envVar))
		}

		fmt.Println("config", config)

		return &config, nil
	}

	fmt.Println("LoadEnv 2")

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	val := reflect.Indirect(reflect.ValueOf(config))
	for i := 0; i < val.NumField(); i++ {
		envVar := val.Type().Field(i).Tag.Get("mapstructure")
		if envVar == "" {
			return nil, fmt.Errorf("Env var '%s' not found", envVar)
		}
	}

	return &config, nil
}
