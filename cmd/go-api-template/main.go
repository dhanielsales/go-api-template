package main

import (
	"fmt"
	"os"
	"runtime/debug"

	_ "github.com/dhanielsales/go-api-template/docs"

	// Set up config
	"github.com/dhanielsales/go-api-template/internal/app"
	"github.com/dhanielsales/go-api-template/internal/config/env"
	"github.com/dhanielsales/go-api-template/internal/config/shutdown"
)

func mainRecover() {
	if err := recover(); err != nil {
		fmt.Printf("panic: %v\n", err)
		debug.PrintStack()
	}
}

// @title Go Template API
// @version 1.0
// @description A simple API to show how to use Go in a clean way
// @contact.name Dhaniel Sales
// @BasePath /
func main() {
	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		fmt.Printf("exiting with code %d\n", exitCode)
		os.Exit(exitCode)
	}()
	defer mainRecover()

	// load config
	envVars := env.GetInstance()

	// Create new service
	srv, err := app.New(envVars)
	if err != nil {
		fmt.Printf("error creating service: %v", err)
		exitCode = 1
		return
	}

	// Start and ensuring the server is shutdown gracefully & app runs
	shutdown.StartGracefully(srv)
}
