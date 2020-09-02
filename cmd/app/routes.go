package main

import (
	"github.com/angelmendozacap/go-structure/database"
	"github.com/angelmendozacap/go-structure/pkg/http/rest/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	connModel := database.Model{
		Database: "snippetbox",
		Engine:   "mysql",
		Port:     3306,
		Password: "kick1930",
		User:     "web",
		Server:   "localhost",
	}

	db, _ := connModel.NewConnection()

	routes.TagRoute(e, "mysql", db)

	return e
}
