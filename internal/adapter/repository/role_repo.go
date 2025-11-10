package repository

import (
	"context"
	"go-app/internal/domain/entity"
	"go-app/internal/domain/repository"
	"go-app/pkg/errors"

	"gorm.io/gorm"
)

// roleRepository ...
type roleRepository struct {
	*gorm.DB
}

// NewRoleRepository will implement of repository.roleRepository interface
func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{
		DB: db,
	}
}

// Fetch will fetch content from db
func (rp *roleRepository) Fetch(ctx context.Context) ([]entity.Role, error) {
	dao := []Role{}
	if err := rp.DB.WithContext(ctx).Find(&dao).Error; err != nil {
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}

	roles := []entity.Role{}

	for i := range dao {
		r := convertRoleToEntity(&dao[i])
		roles = append(roles, *r)
	}

	return roles, nil
}

// Find will find content from db
func (rp *roleRepository) Find(ctx context.Context, id uint) (*entity.Role, error) {
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
func (rp *roleRepository) CheckExists(ctx context.Context, q entity.Role, id *uint) (bool, error) {
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
func (rp *roleRepository) Store(ctx context.Context, role *entity.Role) error {
	dao := convertRoleToDao(role)
	if err := rp.DB.WithContext(ctx).Create(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	*role = *convertRoleToEntity(dao)

	return nil
}

// Update will update data to db
func (rp *roleRepository) Update(ctx context.Context, role *entity.Role) error {
	dao := convertRoleToDao(role)
	if err := rp.DB.WithContext(ctx).Save(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	*role = *convertRoleToEntity(dao)

	return nil
}

// Delete will delete data from db
func (rp *roleRepository) Delete(ctx context.Context, id uint) error {
	dao := Role{}
	if err := rp.DB.WithContext(ctx).Delete(&dao, id).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	return nil
}
