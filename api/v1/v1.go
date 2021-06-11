package v1

import (
	"github.com/labstack/echo"
	"github.com/unico/FeirasLivresAPI/api/v1/feira"
	"github.com/unico/FeirasLivresAPI/app"
)

func Register(g *echo.Group, apps *app.Container) {
	v1 := g.Group("/v1")

	// health.Register(v1.Group("/health"), apps, middleware)
	// mock.Register(v1.Group("/mock"), apps, middleware)
	//feira.Register(v1.Group("/feiras"), apps)

	feira.Register(v1.Group("/feiras"), apps)

}
