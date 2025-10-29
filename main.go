package main

import (
	"cmp"
	"log"
	"os"
	"time"

	_ "spaApp/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// ===== CORS =====
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	ServeReact(r)
	ApiRoutes(r)

	// Swagger документация
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ===== RUN SERVER =====
	port := cmp.Or(os.Getenv("PORT"), "8080")
	ip := cmp.Or(os.Getenv("IP"), "localhost")
	log.Printf("Server starting http://%s:%s", ip, port)
	if err := r.Run(ip + ":" + port); err != nil {
		log.Fatal(err)
	}
}
