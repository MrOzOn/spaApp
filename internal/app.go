package internal

import (
	"embed"
	"log"
	_ "spaApp/docs"
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

	ServeReact(a.r, frontendFS)
	ApiRoutes(a.r)

	// Swagger документация
	a.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return a, nil
}

func (a *App) Start(port, ip string) error {
	log.Printf("Server starting http://%s:%s", ip, port)
	if err := a.r.Run(ip + ":" + port); err != nil {
		return err
	}
	return nil
}
