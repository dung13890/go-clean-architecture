package registry

import (
	"go-app/internal/domain"
	"go-app/internal/impl/service"
	"go-app/pkg/cache"
)

// Service registry
type Service struct {
	JWTSvc      domain.JWTService
	ThrottleSvc domain.ThrottleService
}

// NewService will create new an service object representation of domain.
func NewService(cm cache.Client) *Service {
	return &Service{
		JWTSvc:      service.NewJWTService(cm),
		ThrottleSvc: service.NewThrottleService(cm),
	}
}
