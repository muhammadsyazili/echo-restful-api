package controllers

import (
	"fmt"
	"time"
	"net/http"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/models"
	"github.com/muhammadsyazili/echo-rest/helpers"
	"github.com/muhammadsyazili/echo-rest/template"
)

func CheckLogin(c echo.Context) error {
	var res template.Response

	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.CheckLogin(username, password)

	if err == sql.ErrNoRows {
		fmt.Println("User not found!")

		res.Status = http.StatusBadRequest
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	match, err := helpers.CheckHash(password, result.Password)
	if !match {
		fmt.Println("User not found!")

		res.Status = http.StatusBadRequest
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = result.Id
	claims["username"] = result.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and sendtemplate.response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Status = http.StatusOK
	res.Message = "Ok"
	res.Data = map[string]string{"token": t}
	return c.JSON(http.StatusOK, res)
}