package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/nehemiah-dev/payment-gateway/docs"
	"github.com/nehemiah-dev/payment-gateway/internal/handlers"
	"github.com/nehemiah-dev/payment-gateway/internal/middleware"
)

func New() *gin.Engine {
	r := gin.New() // not gin.Default() — we add our own middleware

	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	r.GET("/health", handlers.Health)
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	_ = v1

	return r
}
