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

func CreateRouter(in *Incoming) *gin.Engine {
	r := gin.Default()

	root := r.Group("/", in.filter2)
	root.GET("/ping", in.pong)

	return r
}

func (in *Incoming) pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
