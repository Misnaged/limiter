package internal

import (
	"fmt"
	"limiter/config"
	"limiter/internal/server"
	"limiter/internal/service"
	"limiter/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	httpsServer *http.Server
	config      *config.Scheme
	srv         *service.Service
}

func (app *App) Serve() error {
	go func() {
		if err := app.httpsServer.ListenAndServe(); err != nil {
			logger.Log().Info(fmt.Errorf("listen HTTP server: %w", err).Error())
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	if err := app.httpsServer.Close(); err != nil {
		return fmt.Errorf("wasn't able to successfully shutdown application: %w", err)
	}
	return nil
}

func (app *App) Init() error {
	var err error
	app.srv, err = service.NewService()
	if err != nil {
		return fmt.Errorf("%w")
	}
	app.httpsServer = server.NewServer(app.config, app.srv)

	return nil
}

func (app *App) Config() *config.Scheme {
	return app.config
}
func NewApplication() *App {
	return &App{
		config: &config.Scheme{},
	}
}
