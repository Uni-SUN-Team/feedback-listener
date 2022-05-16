package controller

import (
	"net/http"
	"unisun/api/feedback-processor-api/src/service"

	"github.com/gin-gonic/gin"
)

func TestGetFeedback(c *gin.Context) {
	result := service.GetFeedbackSumAndCount()
	result.Score = float32(result.Sum / result.Count)
	c.JSON(http.StatusOK, result)
}
