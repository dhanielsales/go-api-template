package app

import (
	"context"
	"database/sql"
	"fmt"

	//nolint:revive // necessary to set up swagger docs.
	_ "github.com/dhanielsales/go-api-template/docs"

	// Set up config
	"github.com/dhanielsales/go-api-template/internal/config/env"

	"github.com/dhanielsales/go-api-template/pkg/httputils"
	"github.com/dhanielsales/go-api-template/pkg/logger"
	"github.com/dhanielsales/go-api-template/pkg/transcriber"

	postgresstorage "github.com/dhanielsales/go-api-template/pkg/postgres"
	redisstorage "github.com/dhanielsales/go-api-template/pkg/redis"

	// Modules
	"github.com/dhanielsales/go-api-template/internal/modules/store"

	goredis "github.com/redis/go-redis/v9"
)

type app struct {
	http      *httputils.HTTPServer
	redis     *redisstorage.Storage
	postgres  *postgresstorage.Storage
	env       *env.Values
	logger    logger.Logger
	transcrib transcriber.Transcriber
}

func New(envVars *env.Values) (*app, error) {
	// init the Postgres storage
	postgresDB, err := sql.Open("postgres", envVars.POSTGRES_URL)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres connection: %w", err)
	}

	postgres := postgresstorage.New(postgresDB)

	// init the Redis storage
	opts, err := goredis.ParseURL(envVars.REDIS_URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis url: %w", err)
	}

	redisClient := goredis.NewClient(opts)
	redisStorage, err := redisstorage.New(redisClient)
	if err != nil {
		return nil, fmt.Errorf("error opening redis connection: %w", err)
	}

	// init logger
	loggerInstance := logger.GetInstance()

	// init http server
	httpServer := httputils.New(envVars)

	// init validator
	transcrib := transcriber.DefaultTranscriber()
	validator := httputils.NewValidator(transcrib)

	// Start store module
	store.Bootstrap(postgres, redisStorage, httpServer, validator)

	return &app{
		http:      httpServer,
		postgres:  postgres,
		redis:     redisStorage,
		logger:    loggerInstance,
		transcrib: transcrib,
		env:       envVars,
	}, nil
}

func (s *app) Run() {
	logger.Info("Starting...")
	s.http.Start()
}

func (s *app) Cleanup(ctx context.Context) error {
	logger.Info("Cleaning up...")
	if err := s.http.Cleanup(ctx); err != nil {
		return err
	}

	if err := s.postgres.Cleanup(); err != nil {
		return err
	}

	if err := s.redis.Cleanup(); err != nil {
		return err
	}

	return nil
}
