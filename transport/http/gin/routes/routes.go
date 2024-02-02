package routes

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/artbasketai/ab-backend/transport/http/gin/handlers"
)

// setupRoutes defines all the rest API endpoints served by this server
func setupRoutes(router *gin.Engine) {
	router.GET("/", handlers.GetRoot)
	// api groups is responsible for serving data only
	api := router.Group("/api")
	{
		// version v1 for later deprecation  &/or replacement
		_ = api.Group("/v1")
		{
			// service related api calls
			// ---------------------------------------------------------------------------------------------------
		}
	}
}

// CreateRouter creates and configures a server
func CreateRouter() *gin.Engine {
	router := gin.New()
	router.Use(
		// Add MiddleWares here if needed
		cors.New(
			cors.Config{
				AbortOnError:     false,
				AllowAllOrigins:  true,
				AllowedOrigins:   nil,
				AllowOriginFunc:  nil, // don't set this
				AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
				AllowedHeaders:   []string{"*"}, // allow everything
				ExposedHeaders:   []string{"*"}, // allow everything very insecure fix later
				AllowCredentials: true,
				MaxAge:           86400,
			},
		),
	)
	setupRoutes(router)
	return router
}
