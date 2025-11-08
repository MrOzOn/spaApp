package internal

import (
	"spaApp/internal/handlers"

	"github.com/gin-gonic/gin"
)

// @title User API
// @version 1.0
// @description Это REST API для управления пользователями
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@myapp.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api
func ApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/time", handlers.GetCurrentTime)
	}
}
