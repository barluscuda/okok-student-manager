package service

import (
	"context"
	"fmt"

	"github.com/okok-student-manager/internal/domain"
	"github.com/okok-student-manager/internal/port"
)

type healthService struct {
	probe port.HealthProbe
}

func NewHealthService(probe port.HealthProbe) *healthService {
	return &healthService{probe: probe}
}

func (s *healthService) Check(ctx context.Context) (domain.HealthStatus, error) {
	if err := s.probe.Check(ctx); err != nil {
		return "", fmt.Errorf("run health probe: %w", err)
	}
	return domain.HealthStatusHealthy, nil
}
