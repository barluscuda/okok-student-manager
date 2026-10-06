package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/okok-student-manager/internal/adapter/http/dto"
	"github.com/okok-student-manager/internal/port"
)

type HealthHandler struct {
	checker port.HealthChecker
}

func NewHealthHandler(checker port.HealthChecker) *HealthHandler {
	return &HealthHandler{checker: checker}
}

func (h *HealthHandler) Check(c echo.Context) error {
	status, err := h.checker.Check(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "service is unhealthy").SetInternal(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Code:    "OK",
		Message: "service is healthy",
		Data: dto.HealthData{
			Status:    string(status),
			Timestamp: time.Now().UTC(),
		},
	})
}
