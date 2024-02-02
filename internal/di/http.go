package di

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/artbasketai/ab-backend/config"
	http_gin "github.com/artbasketai/ab-backend/transport/http/gin"
)

// GetHTTPServer returns a gin implementation for httpServer
func (c *Container) GetHTTPServer() *http.Server {
	if c.cache.hTTPServer != nil {
		return c.cache.hTTPServer
	}

	address := fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort())
	log.Printf("[INFO] Starting http server at address : %s", address)

	gin.SetMode(gin.ReleaseMode)

	server := http_gin.Start(address)


	c.cache.hTTPServer = server

	return c.cache.hTTPServer
}
