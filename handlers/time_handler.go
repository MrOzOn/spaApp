package handlers

import (
	"log"
	"net/http"
	"spaApp/common"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCurrentTime(c *gin.Context) {
	currentTime := time.Now()
	log.Printf("current time is %s", currentTime)
	c.JSON(http.StatusOK, common.APIResponse{Data: currentTime})
}
