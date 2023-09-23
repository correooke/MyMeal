package handler

import (
	"net/http"

	"github.com/correooke/MyMeal/common/model"
	"github.com/correooke/MyMeal/common/service"
	"github.com/labstack/echo/v4"
)

type CommonHandler[T model.Entity] struct {
	Service service.CommonService[T]
}

func NewCommonHandler[T model.Entity](service service.CommonService[T]) CommonHandler[T] {
	return CommonHandler[T]{Service: service}
}

func (h *CommonHandler[T]) GetAll(c echo.Context) error {
	entities, err := h.Service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entities)
}

func (h *CommonHandler[T]) GetByID(c echo.Context) error {
	id := c.Param("id")
	entity, err := h.Service.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entity)
}

func (h *CommonHandler[T]) Create(c echo.Context) error {
	var entity T
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := h.Service.Create(c.Request().Context(), entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//c.Response().Header().Set("Location", fmt.Sprintf("/path/to/resource/%s", entity.ID()))
	return c.JSON(http.StatusCreated, entity)
}

func (h *CommonHandler[T]) Update(c echo.Context) error {
	id := c.Param("id")
	var entity T
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := h.Service.Update(c.Request().Context(), id, entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entity)
}

func (h *CommonHandler[T]) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.Service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
