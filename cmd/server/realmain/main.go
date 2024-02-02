package realmain

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/artbasketai/ab-backend/config"
	"github.com/artbasketai/ab-backend/internal/di"
	"golang.org/x/sync/errgroup"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK = iota
	exitError
)

// Execute is start point of service. We initialize required client
// (DB or external services) and logger and pass it to the server struct.
// The main steps are as follows :
//	1. Read environment variables
//  2. Initialize DI Container
//  3. Inject dependencies
//  4. Inject required servers
//  5. Start listening on separate goroutines
//  6. Handle graceful shutdown
//nolint:funlen // too large and doesn't contain any business logic
func Execute() int {
	_, _ = fmt.Fprint(os.Stdout, "[INFO] Beginning Execution\n")
	//	1. Read environment variables
	// -----------------------------------------------------------------------------------------------------
	env, err := config.ReadFromEnv()
	if err != nil {
		_, _ = fmt.Fprint(os.Stdout, "[ERROR] Reading ENV variables\n")
		return exitError
	}
	//  2. Initialize DI Container
	// -----------------------------------------------------------------------------------------------------
	// Context for container
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, _ = fmt.Fprint(os.Stdout, "[INFO] Creating container\n")
	// Create a DI container
	container := di.NewContainer(ctx, env)
	// 3. Inject dependencies
	// -----------------------------------------------------------------------------------------------------
	// TODO

	// 4. Inject required servers
	// -----------------------------------------------------------------------------------------------------
	// Create http server with the required configurations
	httpServer := container.GetHTTPServer()
	//  5. Start listening on separate goroutines
	// -----------------------------------------------------------------------------------------------------
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return httpServer.ListenAndServe()
	})
	//  6. Handle graceful shutdown
	// -----------------------------------------------------------------------------------------------------
	// Waiting for SIGTERM or Interrupt signal. If server receives them,
	// gRPC server and http server will shutdown gracefully.
	// Waiting for SIGTERM or Interrupt signal. If server receives them,
	// gRPC server and http server will shutdown gracefully.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		_, _ = fmt.Fprint(os.Stdout, "\n[INFO] Received SIGTERM, exiting server gracefully\n")
	case <-ctx.Done():
	}
	// Remember to close all connections like db here
	return exitOK
}
