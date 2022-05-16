package src

import (
	"os"
	"unisun/api/feedback-processor-api/src/route"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(os.Getenv("CONTEXT_PATH") + "/api")
	{
		route.Feedback(g)
	}
	return r
}
