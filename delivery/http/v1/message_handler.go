package v1

import (
	"bytes"
	"io"
	"net/http"

	"github.com/mahfuz110244/project1/entity"
	message "github.com/mahfuz110244/project1/service/message"

	"github.com/labstack/echo/v4"
)

func GetMessageInfoHandler(c echo.Context) error {
	var b bytes.Buffer
	if _, err := io.Copy(&b, c.Request().Body); err != nil {
		c.Logger().Errorf("failed to copy response body: %s", err)
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: "failed to copy response body",
		})

	}
	inputMessage := b.String()
	if inputMessage == "" {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Input message missing!!! please provide valid message as text",
		})
	}
	res, err := message.GetTextInfo(c.Request().Context(), inputMessage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: "Something went wrong!",
		})
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Data:    res,
	})
}
