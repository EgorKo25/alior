package command

import (
	"github.com/EgorKo25/common/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Command interface {
	Name() string
	Apply() error
	Parse(ctx *gin.Context) error
}

type Commander interface {
	GetCommand(name string) gin.HandlerFunc
}

func NewCommand(logger logger.ILogger) (*Manager, error) {
	manager := &Manager{logger: logger}
	manager.Register("callback/create", &CallbackCreate{})

	return manager, nil
}

type Manager struct {
	logger logger.ILogger

	commands map[string]Command
}

func (c *Manager) Register(name string, command Command) {
	if c.commands == nil {
		c.commands = make(map[string]Command)
	}
	c.commands[name] = command
}

func (c *Manager) CastToGin(command Command) gin.HandlerFunc {
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
