package main

import (
	"net/http"

	"github.com/clopst/petstore/internal/petstore/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.GET("/categories", rest.GetAllCategory)
	e.POST("/categories", rest.CreateCategory)
	e.GET("/categories/:id", rest.ShowCategory)

	e.Logger.Fatal(e.Start(":9000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world.")
}
