package main

import (
	"os"
	"unisun/api/feedback-processor-api/src"
	"unisun/api/feedback-processor-api/src/config"
	"unisun/api/feedback-processor-api/src/constants"
)

func main() {
	if os.Getenv(constants.NODE) != constants.PRODUCTION {
		config.ConfigENV()
	}
	config.ConnectDatabase()
	r := src.App()
	port := os.Getenv(constants.PORT)
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
