package app

import (
	"database/sql"
	"fmt"

	_ "github.com/dhanielsales/go-api-template/docs"

	// Set up config
	"github.com/dhanielsales/go-api-template/internal/config/env"

	"github.com/dhanielsales/go-api-template/pkg/httputils"
	"github.com/dhanielsales/go-api-template/pkg/logger"
	"github.com/dhanielsales/go-api-template/pkg/postgres"
	"github.com/dhanielsales/go-api-template/pkg/redis"
	"github.com/dhanielsales/go-api-template/pkg/transcriber"

	// Modules
	"github.com/dhanielsales/go-api-template/internal/modules/store"

	goredis "github.com/redis/go-redis/v9"
)

type app struct {
	http      *httputils.HttpServer
	postgres  *postgres.Storage
	redis     *redis.Storage
	logger    logger.Logger
	env       *env.EnvVars
	transcrib transcriber.Transcriber
}

func New(env *env.EnvVars) (*app, error) {
	// init the Postgres storage
	postgresDb, err := sql.Open("postgres", env.POSTGRES_URL)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres connection: %w", err)
	}

	postgres := postgres.New(postgresDb)

	// init the Redis storage
	opts, err := goredis.ParseURL(env.REDIS_URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis url: %w", err)
	}

	client := goredis.NewClient(opts)
	redis, err := redis.New(client)
	if err != nil {
		return nil, fmt.Errorf("error opening redis connection: %w", err)
	}

	// init logger
	logger := logger.GetInstance()

	// init http server
	httpServer := httputils.New(env)

	// init validator
	transcrib := transcriber.DefaultTranscriber()
	validator := httputils.NewValidator(transcrib)

	// Start store module
	store.Bootstrap(postgres, redis, httpServer, validator)

	return &app{
		http:      httpServer,
		postgres:  postgres,
		redis:     redis,
		logger:    logger,
		transcrib: transcrib,
		env:       env,
	}, nil
}

func (s *app) Run() {
	fmt.Println("Starting...")
	s.http.Start()
}

func (s *app) Cleanup() {
	fmt.Println("Cleaning up...")
	s.http.Cleanup()
	s.postgres.Cleanup()
	s.redis.Cleanup()
}
