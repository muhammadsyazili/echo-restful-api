package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/models"
)


func GetAllAccount(c echo.Context) error {
	var res models.Response

	result, err := models.GetAllAccount()

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func GetWhereAccount(c echo.Context) error {
	var res models.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}
	
	result, err := models.GetWhereAccount(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func StoreAccount(c echo.Context) error {
	var res models.Response

	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.StoreAccount(username, password)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateAccount(c echo.Context) error {
	var res models.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.UpdateAccount(id, username, password)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func DestroyAccount(c echo.Context) error {
	var res models.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := models.DestroyAccount(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}