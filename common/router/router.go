package router

import (
	"fmt"

	"github.com/correooke/MyMeal/common/handler"
	"github.com/correooke/MyMeal/common/model"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	r := echo.New()

	return r
}

func AddIsAlive(r *echo.Echo) *echo.Echo {
	r.GET("/isalive", handler.IsAlive)

	return r
}

func AddCRUDRoutes[T model.Entity](r *echo.Echo, ch *handler.CommonHandler[T], baseRoute string) *echo.Echo {
	r.GET(fmt.Sprintf("/%s", baseRoute), ch.GetAll)
	r.GET(fmt.Sprintf("/%s/{id}", baseRoute), ch.GetByID)
	r.POST(fmt.Sprintf("/%s", baseRoute), ch.Create)
	r.PUT(fmt.Sprintf("/%s/{id}", baseRoute), ch.Update)
	r.DELETE(fmt.Sprintf("/%s/{id}", baseRoute), ch.Delete)

	return r
}
