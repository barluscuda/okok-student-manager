package service

type HealthService interface {
	Check() error
}

type healthService struct{}

func NewHealthService() HealthService {
	return healthService{}
}

func (healthService) Check() error {
	return nil
}
