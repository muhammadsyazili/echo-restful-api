package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/controllers"
	_"github.com/muhammadsyazili/echo-rest/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This is echo!")
	})

	e.GET("/user", controllers.GetAllUser)
	e.GET("/user/:id", controllers.GetWhereUser)
	e.POST("/user", controllers.StoreUser)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DestroyUser)

	e.GET("/transaction", controllers.GetAllTransaction)
	e.GET("/transaction/:id", controllers.GetWhereTransaction)
	e.POST("/transaction", controllers.StoreTransaction)
	e.PUT("/transaction/:id", controllers.UpdateTransaction)
	e.DELETE("/transaction/:id", controllers.DestroyTransaction)

	//e.DELETE("/user/:id", controllers.DestroyUser, middleware.IsAuthenticated)
	
	e.POST("/login", controllers.CheckLogin)
	return e
}