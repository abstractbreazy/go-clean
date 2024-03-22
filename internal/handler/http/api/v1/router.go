package v1

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) AddRoutes(e *echo.Echo) {
	e.GET("/reset", h.Reset)
	e.GET("/", h.Ping)
}
