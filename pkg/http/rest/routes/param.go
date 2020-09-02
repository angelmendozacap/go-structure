package routes

import (
	"database/sql"

	"github.com/angelmendozacap/go-structure/pkg/http/rest/handlers"

	"github.com/labstack/echo/v4"
)

// ParamRoute del paquete tags
func ParamRoute(e *echo.Echo, engine string, db *sql.DB) {
	r := e.Group("/api/v1/params")

	h := handlers.NewParamHandler(engine, db)
	r.POST("", h.Create)
	r.GET("", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.PATCH("/:id/toggle-active", h.ToggleActive)
}
