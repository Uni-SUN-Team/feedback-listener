package ports

import "github.com/gin-gonic/gin"

type ControllerFeedbackPort interface {
	Feedbacks(c *gin.Context)
	FeedbackById(c *gin.Context)
	FeedbackBySlug(c *gin.Context)
}
