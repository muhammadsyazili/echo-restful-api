package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/models"
	"github.com/muhammadsyazili/echo-rest/template"
)


func GetAllUser(c echo.Context) error {
	var res template.Response

	result, err := models.GetAllUser()

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func GetWhereUser(c echo.Context) error {
	var res template.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}
	
	result, err := models.GetWhereUser(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUser(c echo.Context) error {
	var res template.Response

	name := c.FormValue("name")
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.StoreUser(name, username, email, password)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateUser(c echo.Context) error {
	var res template.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	name := c.FormValue("name")
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.UpdateUser(id, name, username, email, password)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func DestroyUser(c echo.Context) error {
	var res template.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := models.DestroyUser(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}