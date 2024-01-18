package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type Starter interface {
	Run()
	Cleanup()
}

func StartGracefully(s Starter) {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	go s.Run()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	s.Cleanup()
}
