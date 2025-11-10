package internal

import (
	"embed"
	"log"
	"spaApp/docs"
	_ "spaApp/docs"
	"spaApp/internal/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	r *gin.Engine
}

func NewApp(frontendFS embed.FS) (*App, error) {
	a := &App{}

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a simple SPA server"
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	//docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	a.r = gin.Default()
	// ===== CORS =====
	a.r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// ===== REACT STATIC =====
	err := a.ServeReact(frontendFS)
	if err != nil {
		return nil, err
	}
	// ===== API =====
	api := a.r.Group("/api")
	{
		api.GET("/time", handlers.GetCurrentTime)
	}
	// ===== SWAGGER =====
	a.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return a, nil
}

func (app *App) Start(port, ip string) error {
	log.Printf("Server starting http://%s:%s", ip, port)
	if err := app.r.Run(ip + ":" + port); err != nil {
		return err
	}
	return nil
}
