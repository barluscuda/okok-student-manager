package port

import (
	"context"

	"github.com/okok-student-manager/internal/domain"
)

// HealthChecker exposes the health use case to delivery adapters.
type HealthChecker interface {
	Check(ctx context.Context) (domain.HealthStatus, error)
}

// HealthProbe abstracts checks against the running process and its dependencies.
type HealthProbe interface {
	Check(ctx context.Context) error
}
