package handlers

import (
	"log"
	"net/http"
	"spaApp/common"
	"time"

	"github.com/gin-gonic/gin"
)

type GetCurrentTimeResponseWrapper struct {
	Data string `json:"data,omitempty" example:"2020-04-06T00:00:00+00:00"`
}

// GetCurrentTime godoc
// @Summary Get current time
// @Description Get current time in default format
// @Tags other
// @Produce json
// @Success 200 {object} common.APIResponse "dfhdslfj"
// @Response 200 {object} GetCurrentTimeResponseWrapper "dddd"
// @Router /api/time [get]
func GetCurrentTime(c *gin.Context) {
	currentTime := time.Now()
	log.Printf("current time is %s", currentTime)
	c.JSON(http.StatusOK, common.APIResponse{Data: currentTime})
}
