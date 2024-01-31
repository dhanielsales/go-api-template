package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Starter interface {
	Run()
	Cleanup()
}

func StartGracefully(s Starter) {
	fmt.Println("StartGracefully")

	quit := make(chan os.Signal, 1)
	defer close(quit)

	go s.Run()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	s.Cleanup()
}
