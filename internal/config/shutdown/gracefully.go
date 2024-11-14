package shutdown

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dhanielsales/go-api-template/pkg/logger"
)

// Starter interface defines the methods required for a component to be started and cleaned up gracefully.
type Starter interface {
	Run(ctx context.Context) error     // Run starts the component and processes in the background.
	Cleanup(ctx context.Context) error // Cleanup shuts down the component gracefully.
}

// SetupGracefully it's a helper to run the [Starter] that was been passed by parameters and ensure they will be stoped gracefully.
func SetupGracefully(ctx context.Context, starter Starter) error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer cancel()

	go starter.Run(ctx)

	select {
	case <-ctx.Done():
		cause := context.Cause(ctx)
		logger.Info("stopping the service", logger.LogErr("cause", cause))
	}

	if err := starter.Cleanup(ctx); err != nil {
		logger.Error("error on cleanup app", logger.LogErr("err", err))
		return fmt.Errorf("error on cleanup app: %w", err)
	}

	return nil
}
