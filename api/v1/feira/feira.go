package feira

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/unico/FeirasLivresAPI/app"
	"github.com/unico/FeirasLivresAPI/model"
)

func Register(g *echo.Group, apps *app.Container) {
	h := &handler{
		apps: apps,
	}

	g.POST("", h.create)
	// 	g.PUT("/:id", h.update)
	// 	g.GET("/:id", h.readOne)
	// 	g.DELETE("/:id", h.delete)
	// 	g.GET("", h.readAll)
}

type handler struct {
	apps *app.Container
}

func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	var request model.Feira
	// if err := c.Bind(&request); err != nil {
	// 	logger.ErrorContext(ctx, "api.v1.feira.create.Bind", err.Error())
	// 	return tcerr.New(http.StatusBadRequest, "Falha ao recuperar dados da requisição", nil)
	// }

	// if err := c.Validate(&request); err != nil {
	// 	logger.ErrorContext(ctx, "api.v1.feira.create.Validate", err.Error())
	// 	return tcerr.New(http.StatusBadRequest, "Requisição Inválida", nil)
	// }

	data, err := h.apps.Feira.Create(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{
		Data: data,
	})
}
