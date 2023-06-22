package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler(echo *echo.Echo,
) {
	handler := &HealthCheckHandler{}
	echo.GET("/health-check", handler.HealthCheck)
	// echo.GET("/status", handler.Status)
}

// CreateTask is method for create task api endpoint
func (h *HealthCheckHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func (h *HealthCheckHandler) Status(c echo.Context) error {
	return c.JSON(http.StatusOK, "testttttttttttt")
}
