package main

import (
	"cmp"
	"embed"
	"log"
	"os"
	"spaApp/internal"
)

//go:embed frontend/dist
var frontendFS embed.FS

func main() {
	port := cmp.Or(os.Getenv("PORT"), "8080")
	ip := cmp.Or(os.Getenv("IP"), "localhost")

	app, err := internal.NewApp(frontendFS)
	if err != nil {
		log.Fatal(err)
	}
	err = app.Start(port, ip)
	if err != nil {
		log.Fatal(err)
	}
}
