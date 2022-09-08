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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func ShowCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect parameter.")
	}

	res, err := service.ShowCategory(id)
	if res == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No data found.")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func CreateCategory(c echo.Context) error {
	category := new(db.Category)

	if err := c.Bind(category); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := service.CreateCategory(category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
