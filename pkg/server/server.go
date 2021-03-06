package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", GetPingPong)

	return r
}

func GetPingPong(c *gin.Context)  {
	c.JSON(200, gin.H{
			"message": "pong",
	})
}