package v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	httpErrs "go-clean/internal/handler/http"
)

func (h *Handler) Reset(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		realIP = c.RealIP()
		err    error
		subnet string
	)

	if subnet, err = h.uc.Reset(ctx, realIP); err != nil {
		code, message := httpErrs.ErrToHTTPStatus(err)
		return c.JSON(code, httpErrs.NewHTTPError(code, string(message)))
	}

	return c.String(http.StatusOK, fmt.Sprintf("Reset rate limit for subnet : %s\n", subnet))
}
