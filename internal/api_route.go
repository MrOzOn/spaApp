package internal

import (
	"spaApp/internal/handlers"
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
func (app *App) ApiRoutes() {
	api := app.r.Group("/api")
	{
		api.GET("/time", handlers.GetCurrentTime)
	}
}
