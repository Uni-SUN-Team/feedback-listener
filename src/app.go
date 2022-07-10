package src

import (
	"strings"
	"unisun/api/feedback-processor-api/docs"
	"unisun/api/feedback-processor-api/src/controllers"
	"unisun/api/feedback-processor-api/src/routes"
	"unisun/api/feedback-processor-api/src/services"
	"unisun/api/feedback-processor-api/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @termsOfService 		http://swagger.io/terms/
// @contact.name 		API Support
// @contact.url 		http://www.swagger.io/support
// @contact.email 		support@swagger.io

// @license.name 		MIT License Copyright (c) 2022 Uni-SUN-Team
// @license.url 		https://api.unisun.dynu.com/strapi-information-gateway/api/license

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func App() *gin.Engine {
	docs.SwaggerInfo.Title = "feedback-listener api"
	docs.SwaggerInfo.Description = "This is a server celler to strapi server."
	docs.SwaggerInfo.Version = viper.GetString("app.version")
	docs.SwaggerInfo.Host = viper.GetString("app.host")
	docs.SwaggerInfo.BasePath = strings.Join([]string{viper.GetString("app.context_path"), viper.GetString("app.root_path")}, "")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	healController := controllers.NewControllerHealthCheckHandler()

	utilsAdap := utils.NewUtilsHTTPRequestAdapter()
	serviceAdap := services.NewConsumerServiceAdapter(utilsAdap)
	controllerAdap := controllers.NewFeedbackController(serviceAdap)
	routeAdap := routes.NewRouteFeedbackAdapter(controllerAdap)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(strings.Join([]string{viper.GetString("app.context_path"), viper.GetString("app.root_path"), "/v1"}, ""))
	{
		g.GET("/healcheck", healController.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		routeAdap.RouteFeedback(g)
	}
	return r
}
