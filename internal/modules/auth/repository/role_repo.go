package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"

	"gorm.io/gorm"
)

// RoleRepository ...
type RoleRepository struct {
	*gorm.DB
}

// NewRoleRepository will implement of domain.RoleRepository interface
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		DB: db,
	}
}

// Fetch will fetch content from db
func (rp *RoleRepository) Fetch(ctx context.Context) ([]domain.Role, error) {
	dao := []Role{}
	if err := rp.DB.WithContext(ctx).Find(&dao).Error; err != nil {
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}

	roles := []domain.Role{}

	for i := range dao {
		r := convertRoleToEntity(&dao[i])
		roles = append(roles, *r)
	}

	return roles, nil
}

// Find will find content from db
func (rp *RoleRepository) Find(ctx context.Context, id int) (*domain.Role, error) {
	dao := Role{}
	if err := rp.DB.WithContext(ctx).First(&dao, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound.Wrap(err)
		}
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}

	return convertRoleToEntity(&dao), nil
}

// CheckExists will check if data is exist or not
func (rp *RoleRepository) CheckExists(ctx context.Context, q domain.Role, id *int) (bool, error) {
	dao := convertRoleToDao(&q)
	var exists bool
	subQuery := rp.DB.WithContext(ctx).
		Model(&Role{}).
		Select("count(*) > 0").
		Where(&dao)

	if id != nil {
		subQuery = subQuery.Where("id <> ?", id)
	}
	if err := subQuery.Find(&exists).Error; err != nil {
		return false, errors.ErrUnexpectedDBError.Wrap(err)
	}

	return exists, nil
}

// Store will create data to db
func (rp *RoleRepository) Store(ctx context.Context, role *domain.Role) error {
	dao := convertRoleToDao(role)
	if err := rp.DB.WithContext(ctx).Create(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	*role = *convertRoleToEntity(dao)

	return nil
}

// Update will update data to db
func (rp *RoleRepository) Update(ctx context.Context, role *domain.Role) error {
	dao := convertRoleToDao(role)
	if err := rp.DB.WithContext(ctx).Save(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	*role = *convertRoleToEntity(dao)

	return nil
}

// Delete will delete data from db
func (rp *RoleRepository) Delete(ctx context.Context, id int) error {
	dao := Role{}
	if err := rp.DB.WithContext(ctx).Delete(&dao, id).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	return nil
}
