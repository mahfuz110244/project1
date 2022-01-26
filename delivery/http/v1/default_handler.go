package v1

import (
	"net/http"

	"github.com/mahfuz110244/project1/config"

	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	Services *config.App
}

func StatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"app_name":    config.APP_NAME,
		"app_version": config.APP_VERSION,
	})
}
