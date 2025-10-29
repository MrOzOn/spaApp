package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/time", func(c *gin.Context) {
			currentTime := time.Now()
			log.Printf("current time is %s", currentTime)
			c.JSON(http.StatusCreated, APIResponse{Data: currentTime})
		})
	}
}
