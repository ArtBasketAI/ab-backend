package di

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/artbasketai/ab-backend/config"
)

// Container is a DI container with all the dependencies listed and methods associated with this struct to fetch those
// dependencies
type Container struct {
	Env *config.Env
	Ctx context.Context

	cache struct {
		logger     *zap.Logger
		hTTPServer *http.Server
		// database 
		// business specific dependencies constructed locally
	}
}

func NewContainer(ctx context.Context, env *config.Env) *Container {
	return &Container{
		Env: env,
		Ctx: ctx,
	}
}
