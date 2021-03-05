package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/models"
)


func GetAllStudent(c echo.Context) error {
	var res models.Response

	result, err := models.GetAllStudent()
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func GetWhereStudent(c echo.Context) error {
	var res models.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}
	
	result, err := models.GetWhereStudent(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}

func StoreStudent(c echo.Context) error {
	var res models.Response

	nim, err := strconv.Atoi(c.Param("nim"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	account_id, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	nama := c.FormValue("nama")
	jurusan := c.FormValue("jurusan")

	result, err := models.StoreStudent(nama, nim, jurusan, account_id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateStudent(c echo.Context) error {
	var res models.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	nim, err := strconv.Atoi(c.Param("nim"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	account_id, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	nama := c.FormValue("nama")
	jurusan := c.FormValue("jurusan")

	result, err := models.UpdateStudent(id, nama, nim, jurusan, account_id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusCreated, result)
}

func DestroyStudent(c echo.Context) error {
	var res models.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()

		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := models.DestroyStudent(id)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, result)
}