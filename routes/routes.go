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

	e.GET("/student", controllers.GetAllStudent, middleware.IsAuthenticated)
	e.GET("/student/:id", controllers.GetWhereStudent, middleware.IsAuthenticated)
	e.POST("/student", controllers.StoreStudent, middleware.IsAuthenticated)
	e.PUT("/student/:id", controllers.UpdateStudent, middleware.IsAuthenticated)
	e.DELETE("/student/:id", controllers.DestroyStudent, middleware.IsAuthenticated)

	e.GET("/account", controllers.GetAllAccount, middleware.IsAuthenticated)
	e.GET("/account/:id", controllers.GetWhereAccount, middleware.IsAuthenticated)
	e.POST("/account", controllers.StoreAccount, middleware.IsAuthenticated)
	e.PUT("/account/:id", controllers.UpdateAccount, middleware.IsAuthenticated)
	e.DELETE("/account/:id", controllers.DestroyAccount, middleware.IsAuthenticated)
	
	e.POST("/login", controllers.CheckLogin)
	return e
}