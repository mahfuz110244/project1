package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mahfuz110244/project1/entity"
)

// Setup all routers
func SetupRouters(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	health := v1.Group("/health")

	v1.GET("/status", StatusHandler)
	v1.GET("/word/occurrence", GetMessageInfoHandler)

	// Login route
	v1.POST("/login", login)

	// Unauthenticated route
	v1.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")

	r.GET("", restricted)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &entity.Response{
			Success: true,
			Message: "Project1 Server is running properly",
		})
	})
}
