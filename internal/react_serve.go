package internal

import (
	"embed"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func getFileSystem(path string, frontendFS embed.FS) (static.ServeFileSystem, error) {
	fs, err := static.EmbedFolder(frontendFS, path)
	return fs, err
}

func (app *App) ServeReact(frontendFS embed.FS) error {
	distFS, err := getFileSystem("frontend/dist", frontendFS)
	if err != nil {
		return err
	}
	app.r.Use(static.Serve("/", distFS))

	app.r.NoRoute(func(c *gin.Context) {
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
	return nil
}
