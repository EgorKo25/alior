package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"API-GW/src/config"
	"API-GW/src/server/command"

	"github.com/EgorKo25/common/logger"
)

// NewServer создаёт новый экземпляр сервера
func NewServer(config *config.ServerConfig, commandManager command.Commander, logger logger.ILogger) *Server {
	return &Server{config, logger, commandManager}
}

type Server struct {
	*config.ServerConfig
	logger         logger.ILogger
	commandManager command.Commander
}

// Run запускает сервер и обеспечивает Graceful Shutdown
func (s *Server) Run(ctx context.Context) error {
	errChan := make(chan error, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	serv := s.getServer()
	go s.listenAndServe(errChan, serv)

	select {
	case sig := <-quit:
		s.logger.Info("received signal: %s", sig)
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("context down with error: %w", err)
		}
		s.logger.Info("context done")
	case err := <-errChan:
		return fmt.Errorf("server down with error: %w", err)
	}

	childCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serv.Shutdown(childCtx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	s.logger.Info("server exited gracefully")
	return nil
}

// getServer создает новый http.Server
func (s *Server) getServer() *http.Server {
	engine := getRouter(s.commandManager)
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Host, s.Port),
		Handler: engine,
	}
}

// listenAndServe запускает сервер и обрабатывает ошибки запуска
func (s *Server) listenAndServe(errChan chan error, server *http.Server) {
	if s.TlsCert != "" {
		if err := server.ListenAndServeTLS(s.TlsCert, s.KeyFile); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	} else {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}
	close(errChan)
}
