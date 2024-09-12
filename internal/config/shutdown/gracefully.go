package shutdown

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dhanielsales/go-api-template/pkg/logger"
)

type Starter interface {
	Run()
	Cleanup() error
}

func StartGracefully(s Starter) {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	go s.Run()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	if err := s.Cleanup(); err != nil {
		logger.Error("error on cleanup app", logger.LogErr("err", err))
	}
}
