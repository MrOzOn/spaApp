package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist
var frontendFS embed.FS

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func getFileSystem(path string) static.ServeFileSystem {
	fs, err := static.EmbedFolder(frontendFS, path)
	if err != nil {
		log.Fatal(err)
	}
	return fs
}

func Serve(app *gin.Engine) {
	distFS := getFileSystem("frontend/dist")
	app.Use(static.Serve("/", distFS))

	app.NoRoute(func(c *gin.Context) {
		// Only serve index.html for non-API routes
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			index, err := distFS.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer index.Close()
			stat, _ := index.Stat()
			http.ServeContent(c.Writer, c.Request, "index.html", stat.ModTime(), index)
		}
	})
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

	Serve(r)

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
