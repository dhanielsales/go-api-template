package main

import (
	"fmt"
	"os"

	"github.com/dhanielsales/golang-scaffold/config"
	_ "github.com/dhanielsales/golang-scaffold/docs"
	"github.com/dhanielsales/golang-scaffold/init/shutdown"
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/internal/mongo"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"

	"github.com/dhanielsales/golang-scaffold/modules/category"
)

type service struct {
	http     *http.HttpServer
	mongodb  *mongo.Storage
	postgres *postgres.Storage
	env      config.EnvVars
}

func new(env config.EnvVars) (*service, error) {
	// init the Mongodb storage
	// mongodb, err := mongo.Bootstrap(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	// if err != nil {
	// 	return nil, err
	// }

	// init the Postgres storage
	postgres, err := postgres.Bootstrap(env.POSTGRES_URI)
	if err != nil {
		return nil, err
	}

	// init http server
	httpServer := http.Bootstrap(env.PORT)

	// Start Category module
	category.Bootstrap(postgres, httpServer)

	return &service{
		http: httpServer,
		// mongodb:  mongodb,
		postgres: postgres,
		env:      env,
	}, nil
}

func (s *service) run() {
	s.http.Start()
}

func (s *service) cleanup() {
	s.http.Cleanup()
	s.mongodb.Cleanup()
	s.postgres.Cleanup()
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
	env, err := config.LoadConfig()
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

	// run the cleanup after the server is terminated
	defer srv.cleanup()

	// run the server
	go srv.run()

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}
