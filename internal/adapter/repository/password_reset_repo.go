package repository

import (
	"context"
	"time"

	"go-app/internal/domain/repository"
	"go-app/internal/infrastructure/constant"
	"go-app/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// passwordResetRepository ...
type passwordResetRepository struct {
	*gorm.DB
}

// NewPasswordResetRepository will implement of domain.passwordResetRepository interface
func NewPasswordResetRepository(db *gorm.DB) repository.PasswordResetRepository {
	return &passwordResetRepository{
		DB: db,
	}
}

// StoreOrUpdate will store or update password reset by email
func (rp *passwordResetRepository) StoreOrUpdate(ctx context.Context, email, token string) error {
	dao := &PasswordReset{
		Email: email,
		Token: token,
	}

	if err := rp.DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"token", "created_at", "updated_at"}),
	}).Create(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	return nil
}

// FindEmailByToken will find password reset by token
func (rp *passwordResetRepository) FindEmailByToken(ctx context.Context, token string) (string, error) {
	dao := &PasswordReset{
		Token: token,
	}

	createdAt := time.Now().Add(-constant.TokenResetPasswordLifetime)

	if err := rp.DB.WithContext(ctx).Where("created_at >= ?", createdAt).First(&dao, &dao).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.ErrAuthInvalidateToken.Wrap(err)
		}
		return "", errors.ErrUnexpectedDBError.Wrap(err)
	}

	return dao.Email, nil
}

// Delete will delete password reset by token
func (rp *passwordResetRepository) Delete(ctx context.Context, email, token string) error {
	dao := &PasswordReset{
		Email: email,
		Token: token,
	}

	if err := rp.DB.WithContext(ctx).Delete(&dao, &dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	return nil
}
