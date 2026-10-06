package domain

// HealthStatus is the domain-level result of the application health check.
type HealthStatus string

const (
	HealthStatusHealthy HealthStatus = "ok"
)
