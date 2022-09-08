package rest

import (
	"net/http"
	"strconv"

	db "github.com/clopst/petstore/internal/petstore/repo/postgres"
	"github.com/clopst/petstore/internal/petstore/service"
	"github.com/labstack/echo/v4"
)

func GetAllCategory(c echo.Context) error {
	res, err := service.GetAllCategory()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, res)
}

func ShowCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := service.ShowCategory(id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, res)
}

func CreateCategory(c echo.Context) error {
	category := new(db.Category)

	if err := c.Bind(category); err != nil {
		panic(err)
	}

	res, err := service.CreateCategory(category)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, res)
}
