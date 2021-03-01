package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadsyazili/echo-rest/models"
)


func GetAllStudent(c echo.Context) error {
	fmt.Println("sampe controller")
	result, err := models.GetAllStudent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetWhereStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	result, err := models.GetWhereStudent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreStudent(c echo.Context) error {
	nim, err := strconv.Atoi(c.Param("nim"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	nama := c.FormValue("nama")
	jurusan := c.FormValue("jurusan")

	result, err := models.StoreStudent(nama, nim, jurusan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	nim, err := strconv.Atoi(c.Param("nim"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	nama := c.FormValue("nama")
	jurusan := c.FormValue("jurusan")

	result, err := models.UpdateStudent(id, nama, nim, jurusan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DestroyStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := models.DestroyStudent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}