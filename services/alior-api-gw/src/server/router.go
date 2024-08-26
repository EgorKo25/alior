package server

import (
	"API-GW/src/server/command"

	"github.com/gin-gonic/gin"
)

func getRouter(commands command.Commander) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	var c command.ICommand
	c = &command.CallbackCreate{}
	v1.POST("/"+c.Name(), commands.Register(c))
	c = &command.CarouselList{}
	v1.GET("/"+c.Name(), commands.Register(c))
	return router
}
