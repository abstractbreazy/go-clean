package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	httpErrs "go-clean/internal/handler/http"
)

func (h *Handler) Ping(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		ip  = c.RealIP()
		err error
	)

	if err = h.uc.Ping(ctx, ip); err != nil {
		code, message := httpErrs.ErrToHTTPStatus(err)
		return c.JSON(code, httpErrs.NewHTTPError(code, string(message)))
	}

	return c.String(http.StatusOK, "Pong\n")
}
