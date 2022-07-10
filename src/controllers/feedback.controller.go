package controllers

import (
	"unisun/api/feedback-processor-api/src/ports"

	"github.com/gin-gonic/gin"
)

type FeedbackController struct {
	Service ports.ServicesFeedbackPort
}

func NewFeedbackController(service ports.ServicesFeedbackPort) *FeedbackController {
	return &FeedbackController{
		Service: service,
	}
}

func (srv *FeedbackController) Feedbacks(c *gin.Context) {
	// requestPayload := models.ServiceIncomeRequest{}

}

func (srv *FeedbackController) FeedbackById(c *gin.Context) {

}

func (srv *FeedbackController) FeedbackBySlug(c *gin.Context) {

}
