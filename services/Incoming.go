package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Incoming struct {
}

func CreateService() *Incoming {
	s := &Incoming{}
	return s
}

func CreateRouter(myservice *Incoming) *gin.Engine {
	r := gin.Default()

	root := r.Group("/")
	root.GET("/ping", pong)

	return r
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
