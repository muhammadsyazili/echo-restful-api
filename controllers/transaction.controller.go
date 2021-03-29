package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/models"
	"github.com/muhammadsyazili/echo-rest/template"
)


func GetAllTransaction(c echo.Context) error {
	var res template.Response

	result, err := models.GetAllTransaction()
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func GetWhereTransaction(c echo.Context) error {
	var res template.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}
	
	result, err := models.GetWhereTransaction(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func StoreTransaction(c echo.Context) error {
	var res template.Response

	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	amount, err := strconv.ParseFloat(c.Param("amount"), 10)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	title := c.FormValue("title")
	time := c.FormValue("time")
	type_ := c.FormValue("type")

	result, err := models.StoreTransaction(user_id, title, amount, time, type_)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateTransaction(c echo.Context) error {
	var res template.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	amount, err := strconv.ParseFloat(c.Param("amount"), 10)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	title := c.FormValue("title")
	time := c.FormValue("time")
	type_ := c.FormValue("type")

	result, err := models.UpdateTransaction(id, user_id, title, amount, time, type_)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func DestroyTransaction(c echo.Context) error {
	var res template.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := models.DestroyTransaction(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}