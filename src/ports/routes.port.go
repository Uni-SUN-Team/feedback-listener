package ports

import "github.com/gin-gonic/gin"

type RouteFeedback interface {
	RouteFeedback(g *gin.RouterGroup)
}
