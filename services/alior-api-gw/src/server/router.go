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
	v1.GET("/carousel/list", commands.GetCommand("carousel/list"))
	v1.GET("/service/list", commands.GetCommand("service/list"))
	v1.POST("/callback/create", commands.GetCommand("callback/create"))
}
