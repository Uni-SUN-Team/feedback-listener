package main

import (
	"log"
	"unisun/api/feedback-processor-api/src"
	"unisun/api/feedback-processor-api/src/configs"

	"github.com/spf13/viper"
)

func main() {
	envService := configs.New("app", "./src/resource")
	if err := envService.ConfigENV(); err != nil {
		log.Panic(err)
	}
	r := src.App()
	port := viper.GetString("app.port")
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
