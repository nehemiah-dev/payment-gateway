package handlers

import (
	"net/http"

	"github.com/nehemiah-dev/payment-gateway/internal/api"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// GetHealth handles GET /health
func (s *Server) GetHealth(c *gin.Context) {
	status := api.HealthStatusHealthy.Ptr()
	c.JSON(http.StatusOK, api.GetHealth200JSONResponse{Status: status})
}
