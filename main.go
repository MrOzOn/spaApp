package main

import (
	"embed"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var frontendFS embed.FS

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

	// ===== API =====
	api := r.Group("/api")
	{
		api.GET("/time", func(c *gin.Context) {
			currentTime := time.Now()
			log.Printf("current time is %s", currentTime)
			c.JSON(http.StatusCreated, APIResponse{Data: currentTime})
		})
	}

	// ===== React build FS =====
	reactDist, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	// ===== NoRoute для SPA и статики =====
	r.NoRoute(func(c *gin.Context) {
		reqPath := c.Request.URL.Path
		if strings.Contains(reqPath, "..") {
			c.Status(http.StatusBadRequest)
			return
		}

		// Убираем ведущий слэш
		filePath := strings.TrimPrefix(reqPath, "/")
		if filePath != "" {
			// Проверяем наличие файла в embed
			f, err := reactDist.Open(filePath)
			if err == nil {
				f.Close()
				// Определяем MIME
				c.Writer.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(filePath)))
				c.FileFromFS(filePath, http.FS(reactDist))
				return
			}
		}

		// Если файл не найден — отдаём SPA index.html
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.FileFromFS("index.html", http.FS(reactDist))
	})

	// ===== Запуск =====
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
