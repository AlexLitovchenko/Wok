package main

import (
	handler "ServiceB/Handler"

	"github.com/gin-gonic/gin"
)

func main() {
	//handler.Connection()
	router := gin.Default()
	router.GET("/A", handler.GetAll)
	router.Run(":8081")
}
