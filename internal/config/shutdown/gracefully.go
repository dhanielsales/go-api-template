package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dhanielsales/go-api-template/pkg/logger"
)

type Starter interface {
	Run()
	Cleanup(ctx context.Context) error
}

func StartGracefully(s Starter) {
	ctx := context.Background()
	quit := make(chan os.Signal, 1)
	defer close(quit)

	go s.Run()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	if err := s.Cleanup(ctx); err != nil {
		logger.Error("error on cleanup app", logger.LogErr("err", err))
	}
}
