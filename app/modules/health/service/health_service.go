package health_service

type HealthService struct{}

type IHealthService interface {
	Index() error
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (h *HealthService) Index() string {

	return "OK"
}
