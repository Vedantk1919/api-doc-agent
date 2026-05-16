package main

import (
	"os"
	"sample_api_go/nsdConfig"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	nsdConfig.SetupRouter(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
