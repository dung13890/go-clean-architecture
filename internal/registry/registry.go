package registry

import (
	"go-app/internal/adapter/cache"
	"go-app/internal/adapter/mail"
	"go-app/internal/adapter/repository"
	dmSvc "go-app/internal/domain/service"
	"go-app/internal/service"
	"go-app/internal/usecase"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Registry struct
type Registry struct {
	AuthUc *usecase.AuthUsecase
	UserUc *usecase.UserUsecase
	RoleUc *usecase.RoleUsecase
	JWTSvc dmSvc.JWTService
}

// NewRegistry will create new registry
func NewRegistry(db *gorm.DB, rdb *redis.Client) *Registry {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	passwordResetRepo := repository.NewPasswordResetRepository(db)

	cm := cache.NewRedisStore(rdb)
	mailSvc := mail.NewEmail()
	// Initialize services with repositories
	jwtSvc := service.NewJWTService(cm)
	throttleSvc := service.NewThrottleService(cm)

	return &Registry{
		AuthUc: usecase.NewAuthUsecase(jwtSvc, throttleSvc, mailSvc, userRepo, passwordResetRepo),
		UserUc: usecase.NewUserUsecase(userRepo),
		RoleUc: usecase.NewRoleUsecase(roleRepo),
		JWTSvc: jwtSvc,
	}
}
