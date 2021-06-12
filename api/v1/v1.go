package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/unico/FeirasLivresAPI/api/v1/feira"
	"github.com/unico/FeirasLivresAPI/app"
)

func Register(g *echo.Group, apps *app.Container) {
	v1 := g.Group("/v1")

	feira.Register(v1.Group("/feiras"), apps)

}
