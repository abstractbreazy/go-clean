package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerRouter interface {
	AddRoutes(r *echo.Echo)
}

type Router struct {
	*echo.Echo
}

func NewRouter(e *echo.Echo) *Router {
	return &Router{Echo: e}
}

func (r *Router) WithHandler(h HandlerRouter) *Router {
	r.Echo.Use(middleware.Logger())
	r.Echo.Use(middleware.Recover())
	h.AddRoutes(r.Echo)
	return r
}
