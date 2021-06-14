package feira

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/unico/FeirasLivresAPI/app"
	"github.com/unico/FeirasLivresAPI/logger"
	"github.com/unico/FeirasLivresAPI/model"
)

func Register(g *echo.Group, apps *app.Container) {
	h := &handler{
		apps: apps,
	}

	g.POST("", h.create)
	g.PUT("/:id", h.update)
	g.GET("/:id", h.readOne)
	g.DELETE("/:id", h.delete)
	g.GET("/search", h.search)
}

type handler struct {
	apps *app.Container
}

func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	var request model.FeiraRequest
	if err := c.Bind(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.feira.create.Bind", err.Error())
		logger.L.Println(time.Now(), "app.feira.create.Bind", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Falha ao recuperar dados da requisição",
		})
	}

	if err := c.Validate(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.feira.create.Validate", err.Error())
		logger.L.Println(time.Now(), "app.feira.create.Validate", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Requisição Inválida",
		})
	}

	data, err := h.apps.Feira.Create(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{
		Data: data,
	})
}

func (h *handler) update(c echo.Context) error {
	ctx := c.Request().Context()

	var request model.FeiraRequest
	if err := c.Bind(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.feira.update.Bind", err.Error())
		logger.L.Println(time.Now(), "app.feira.update.Bind", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Falha ao recuperar dados da requisição",
		})
	}

	if err := c.Validate(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.feira.update.Validate", err.Error())
		logger.L.Println(time.Now(), "app.feira.update.Validate", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Requisição Inválida",
		})
	}

	id := c.Param("id")
	if id == "" {
		logger.ErrorContext(ctx, "api.v1.feira.update", "o campo 'id' é obrigatório")
		logger.L.Println(time.Now(), "api.v1.feira.update", "o campo 'id' é obrigatório")
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Requisição Inválida",
		})
	}

	data, err := h.apps.Feira.Update(ctx, id, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: data,
	})
}

func (h *handler) readOne(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		logger.ErrorContext(ctx, "api.v1.feira.readOne", "o campo 'id' é obrigatório")
		logger.L.Println(time.Now(), "api.v1.feira.readOne", "o campo 'id' é obrigatório")
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Falha ao recuperar dados da requisição",
		})
	}

	data, err := h.apps.Feira.ReadOne(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: data,
	})
}

func (h *handler) delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		logger.ErrorContext(ctx, "api.v1.feira.delete", "o campo 'id' é obrigatório")
		logger.L.Println(time.Now(), "api.v1.feira.delete", "o campo 'id' é obrigatório")
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Requisição Inválida",
		})
	}

	err := h.apps.Feira.Delete(ctx, id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *handler) search(c echo.Context) error {
	ctx := c.Request().Context()

	var request searchFeira
	if err := c.Bind(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.feira.search.Bind", err.Error())
		logger.L.Println(time.Now(), "app.feira.search.Bind", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Falha ao recuperar dados da requisição",
		})
	}

	if err := c.Validate(&request); err != nil {
		logger.ErrorContext(ctx, "api.v1.feira.search.Validate", err.Error())
		logger.L.Println(time.Now(), "api.v1.feira.search.Validate", err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: "Requisição Inválida",
		})
	}

	data, err := h.apps.Feira.Search(ctx, request.Distrito)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: data,
	})
}
