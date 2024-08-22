package server

import (
	"API-GW/src/server/command"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen ... TODO
type Handlers interface {
	CarouselList(ctx *gin.Context)
	ServiceList(ctx *gin.Context)
	CallbackCreate(ctx *gin.Context)
}

func GetRouter(commands command.Commander) {
	router := gin.Default()
	v1 := router.Group("/v1")
	var c command.ICommand
	c = &command.CallbackCreate{}
	v1.POST("/"+c.Name(), commands.Register(c))
}
