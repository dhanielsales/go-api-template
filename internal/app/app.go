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

	redisstorage "github.com/dhanielsales/go-api-template/pkg/redisutils"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	// Modules
	"github.com/dhanielsales/go-api-template/internal/modules/store"

	_ "github.com/lib/pq"
	goredis "github.com/redis/go-redis/v9"
)

type app struct {
	http      *httputils.HTTPServer
	redis     *redisstorage.Storage
	sql       *sqlutils.Storage
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

	sql := sqlutils.New(postgresDB)
	logger.Info("postgres connection stablished")

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
	logger.Info("redis connection stablished")

	// init logger
	loggerInstance := logger.GetInstance()

	// init http server
	httpServer := httputils.New(envVars)

	// init validator
	transcrib := transcriber.DefaultTranscriber()
	validator := httputils.NewValidator(transcrib)

	// Start store module
	store.Bootstrap(sql, redisStorage, httpServer, validator)

	return &app{
		http:      httpServer,
		sql:       sql,
		redis:     redisStorage,
		logger:    loggerInstance,
		transcrib: transcrib,
		env:       envVars,
	}, nil
}

func (a *app) Run(_ context.Context) error {
	logger.Info(fmt.Sprintf("http server start at %s port", a.env.HTTP_PORT))
	if err := a.http.Start(); err != nil {
		return err
	}

	return nil
}

func (s *app) Cleanup(ctx context.Context) error {
	logger.Info("cleaning up...")
	if err := s.http.Cleanup(ctx); err != nil {
		return err
	}

	if err := s.sql.Cleanup(); err != nil {
		return err
	}

	if err := s.redis.Cleanup(); err != nil {
		return err
	}

	return nil
}
