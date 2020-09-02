package routes

import (
	"database/sql"

	"github.com/angelmendozacap/go-structure/pkg/http/rest/handlers"

	"github.com/labstack/echo/v4"
)

// TagRoute del paquete tags
func TagRoute(e *echo.Echo, engine string, db *sql.DB) {
	r := e.Group("/api/v1/tags")

	h := handlers.NewTagHandler(engine, db)
	r.POST("", h.Create)
	r.GET("", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
