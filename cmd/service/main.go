package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"

	_ "github.com/dhanielsales/golang-scaffold/docs"

	// Config
	"github.com/dhanielsales/golang-scaffold/config/env"
	"github.com/dhanielsales/golang-scaffold/config/log"
	"github.com/dhanielsales/golang-scaffold/config/shutdown"

	// Internal
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/internal/redis"

	// Modules
	"github.com/dhanielsales/golang-scaffold/modules/store"
)

type service struct {
	http     *http.HttpServer
	postgres *postgres.Storage
	redis    *redis.Storage
	logger   log.Logger
	env      *env.EnvVars
	validate *validator.Validate
}

func new(env *env.EnvVars) (*service, error) {
	// init the Postgres storage
	postgres, err := postgres.Bootstrap(env.POSTGRES_URL)
	if err != nil {
		return nil, err
	}

	// init the Redis storage
	redis, err := redis.Bootstrap(env.REDIS_URL)
	if err != nil {
		return nil, err
	}

	// init logger
	logger := log.New(env.APP_NAME)

	// init http server
	httpServer := http.Bootstrap(env.PORT, logger)

	// init validator
	validate := validator.New(validator.WithRequiredStructEnabled())
	validator := http.NewValidator(validate)

	// Start store module
	store.Bootstrap(postgres, redis, httpServer, validator)

	return &service{
		http:     httpServer,
		postgres: postgres,
		redis:    redis,
		logger:   logger,
		validate: validate,
		env:      env,
	}, nil
}

func (s *service) Run() {
	s.http.Start()
}

func (s *service) Cleanup() {
	fmt.Println("Cleaning up...")
	s.http.Cleanup()
	s.postgres.Cleanup()
	s.redis.Cleanup()
}

// @title Golang scaffold
// @version 1.0
// @description A simple Golang backend scaffold
// @contact.name Dhaniel Sales
// @license.name MIT
// @BasePath /
func main() {
	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := env.LoadEnv()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// Create new service
	srv, err := new(env)
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// Start and ensuring the server is shutdown gracefully & app runs
	shutdown.StartGracefully(srv)
}
