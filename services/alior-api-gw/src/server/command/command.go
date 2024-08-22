package command

import (
	"github.com/EgorKo25/common/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ICommand interface {
	Name() string
	Apply() error
	Parse(ctx *gin.Context) error
}

type Commander interface {
	// GetCommand(name string) gin.HandlerFunc
	Register(command ICommand) gin.HandlerFunc
}

func NewCommand(logger logger.ILogger) (*Manager, error) {
	manager := &Manager{logger: logger}
	manager.Register(&CallbackCreate{})

	return manager, nil
}

type Manager struct {
	logger logger.ILogger

	commands map[string]ICommand
}

func (c *Manager) Register(command ICommand) gin.HandlerFunc {
	if c.commands == nil {
		c.commands = make(map[string]ICommand)
	}
	c.commands[command.Name()] = command
	return func(ctx *gin.Context) {
		c.logger.Info("New request to command: %s", command.Name())
		if err := command.Parse(ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := command.Apply(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
