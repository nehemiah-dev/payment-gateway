package handlers

import (
	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Health check
// @Tags system
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
