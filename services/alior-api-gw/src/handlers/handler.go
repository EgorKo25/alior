package handlers

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Parse(c *gin.Context)
}
