package main

import (

	"github.com/AlexLitovchenko/Wok/ServiceB/Handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/A", handler.GetAll)
	router.Run(":8081")
}
