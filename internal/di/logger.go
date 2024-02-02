package di

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/artbasketai/ab-backend/pkg/log"
)

// GetLogger created zap Logger with the loglevel retrieved from the env config
func (c *Container) GetLogger() *zap.Logger {
	if c.cache.logger != nil {
		return c.cache.logger
	}
	logger, err := log.New(c.Env.LogLevel)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "[ERROR] Initializing logger")
		panic(err)
	}

	c.cache.logger = logger

	return c.cache.logger
}
