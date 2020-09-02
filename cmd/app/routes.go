package main

import (
	"github.com/angelmendozacap/go-structure/pkg/http/rest/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := newConnection()

	routes.ParamRoute(e, "mysql", db)

	return e
}
