package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (in *Incoming) filter(ctx *gin.Context) {
	log.Println("test")

	//ctx.JSON(http.StatusOK, "I'm your Filter!")
}
