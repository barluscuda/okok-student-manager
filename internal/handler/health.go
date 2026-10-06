package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/okok-student-manager/internal/handler/dto"
)

type healthChecker interface {
	Check() error
}

type HealthHandler struct {
	service healthChecker
}

func NewHealthHandler(service healthChecker) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) Check(c echo.Context) error {
	if err := h.service.Check(); err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "service is unhealthy").SetInternal(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Code:    "OK",
		Message: "service is healthy",
		Data: dto.HealthData{
			Status:    "ok",
			Timestamp: time.Now().UTC(),
		},
	})
}
