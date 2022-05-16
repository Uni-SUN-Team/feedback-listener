package route

import (
	"unisun/api/feedback-processor-api/src/controller"

	"github.com/gin-gonic/gin"
)

func Feedback(g *gin.RouterGroup) {
	g.GET("/feedbacks", controller.TestGetFeedback)
}
