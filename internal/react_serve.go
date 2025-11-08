package internal

import (
	"embed"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func getFileSystem(path string, frontendFS embed.FS) static.ServeFileSystem {
	fs, err := static.EmbedFolder(frontendFS, path)
	if err != nil {
		log.Fatal(err)
	}
	return fs
}

func ServeReact(app *gin.Engine, frontendFS embed.FS) {
	distFS := getFileSystem("frontend/dist", frontendFS)
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
