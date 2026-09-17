package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/nehemiah-dev/payment-gateway/internal/api"
	"github.com/nehemiah-dev/payment-gateway/internal/handlers"
	"github.com/nehemiah-dev/payment-gateway/internal/middleware"
)

func New() *gin.Engine {
	r := gin.New() // we add our own middleware

	r.Use(middleware.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"X-Request-ID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())

	srv := handlers.NewServer()

	r.GET("/health", srv.GetHealth)
	api.RegisterDocsRoutes(r)

	v1 := r.Group("/api/v1")

	v1.GET("/health", srv.GetHealth)

	return r
}
