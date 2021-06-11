package api

import (
	"github.com/labstack/echo"
	"github.com/micro/go-micro/logger"

	v1 "github.com/unico/FeirasLivresAPI/api/v1"
	"github.com/unico/FeirasLivresAPI/app"
)

// Options struct de opções para a criação de uma instancia das rotas
type Options struct {
	Group *echo.Group
	Apps  *app.Container
}

// Register api instance
func Register(opts Options) {
	v1.Register(opts.Group, opts.Apps)

	logger.Info("Registered -> Api")
}
