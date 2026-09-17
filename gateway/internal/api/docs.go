package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterDocsRoutes(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Ficmart Payment Gateway API - Swagger UI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = () => {
      SwaggerUIBundle({
        url: '/docs/openapi',
        dom_id: '#swagger-ui',
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
        layout: 'StandaloneLayout'
      });
    };
  </script>
</body>
</html>`)
	})

	r.GET("/docs/openapi", func(c *gin.Context) {
		spec, err := GetSwagger()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load OpenAPI spec"})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, spec)
	})
}
