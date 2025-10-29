package main

import (
	"spaApp/handlers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/time", handlers.GetCurrentTime)
	}
}
