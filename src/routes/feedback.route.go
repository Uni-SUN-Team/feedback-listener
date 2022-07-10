package routes

import (
	"unisun/api/feedback-processor-api/src/ports"

	"github.com/gin-gonic/gin"
)

type RouteFeedbackAdapter struct {
	Controller ports.ControllerFeedbackPort
}

func NewRouteFeedbackAdapter(controller ports.ControllerFeedbackPort) *RouteFeedbackAdapter {
	return &RouteFeedbackAdapter{
		Controller: controller,
	}
}

func (srv *RouteFeedbackAdapter) RouteFeedback(g *gin.RouterGroup) {
	g.GET("/feedbacks", srv.Controller.Feedbacks)
	g.GET("/feedbacks/:id", srv.Controller.FeedbackById)
	g.GET("/feedbacks/slug/:slug", srv.Controller.FeedbackBySlug)
}
