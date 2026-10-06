package process

import "context"

// HealthProbe reports whether the API process is still able to serve work.
// Dependency-specific checks can be added here as the application gains them.
type HealthProbe struct{}

func NewHealthProbe() *HealthProbe {
	return &HealthProbe{}
}

func (*HealthProbe) Check(ctx context.Context) error {
	return ctx.Err()
}
