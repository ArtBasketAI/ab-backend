package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoot prints a json content when the root address is clicked
func GetRoot(c *gin.Context) {
	d := []string{
		"You have reached the home page",
	}
	c.JSON(http.StatusOK, d)
}
