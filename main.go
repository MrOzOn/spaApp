package main

import (
	"cmp"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

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

	// ===== RUN SERVER =====
	port := cmp.Or(os.Getenv("PORT"), "8080")
	ip := cmp.Or(os.Getenv("IP"), "localhost")
	log.Printf("Server starting http://%s:%s", ip, port)
	if err := r.Run(ip + ":" + port); err != nil {
		log.Fatal(err)
	}
}
