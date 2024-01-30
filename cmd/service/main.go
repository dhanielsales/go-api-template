package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	goredis "github.com/redis/go-redis/v9"

	_ "github.com/dhanielsales/golang-scaffold/docs"

	// Config
	"github.com/dhanielsales/golang-scaffold/config/env"
	"github.com/dhanielsales/golang-scaffold/config/log"
	"github.com/dhanielsales/golang-scaffold/config/shutdown"

	// Internal
	"github.com/dhanielsales/golang-scaffold/internal/gql"
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/internal/redis"

	// Modules
	"github.com/dhanielsales/golang-scaffold/modules/store"
)

type service struct {
	http        *http.HttpServer
	postgres    *postgres.Storage
	redis       *redis.Storage
	logger      log.Logger
	clientIdeal *gql.Client
	env         *env.EnvVars
	validate    *validator.Validate
}

func new(env *env.EnvVars) (*service, error) {
	// init the Postgres storage
	postgresDb, err := sql.Open("postgres", env.POSTGRES_URL)
	if err != nil {
		return nil, err
	}

	postgres := postgres.Bootstrap(postgresDb)

	// init the Redis storage
	opts, err := goredis.ParseURL(env.REDIS_URL)
	if err != nil {
		return nil, err
	}

	client := goredis.NewClient(opts)
	redis, err := redis.Bootstrap(client)
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

	// init ideal client
	clientIdeal := gql.NewClient(env.EXTERNAL_URL, nil)

	// Start store module
	store.Bootstrap(postgres, redis, clientIdeal, httpServer, validator)

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

// @title Go Scaffold API
// @version 1.0
// @description A simple API to show how to use Go in a clean way
// @contact.name Dhaniel Sales
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
