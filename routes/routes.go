package routes

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/controllers"
	"github.com/muhammadsyazili/echo-rest/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This is echo!")
	})

	e.GET("/user", controllers.GetAllUser, middleware.IsAuthenticated)
	e.GET("/user/:id", controllers.GetWhereUser, middleware.IsAuthenticated)
	e.POST("/user", controllers.StoreUser, middleware.IsAuthenticated)
	e.PUT("/user/:id", controllers.UpdateUser, middleware.IsAuthenticated)
	e.DELETE("/user/:id", controllers.DestroyUser, middleware.IsAuthenticated)
	
	e.POST("/login", controllers.CheckLogin)
	return e
}