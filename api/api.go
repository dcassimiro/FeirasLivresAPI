package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
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

	fmt.Println("Registered -> Api")
}
