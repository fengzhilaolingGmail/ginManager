package repository

import (
	"context"
	"ginManager/models/entity"

	"gorm.io/gorm"
)

type RoleRepo struct{}

func NewRoleRepo() *RoleRepo { return &RoleRepo{} }

// Page 分页
func (r *RoleRepo) Page(ctx context.Context, name string, page, limit int) (list []entity.Role, total int64, err error) {
	db := DB.WithContext(ctx).Model(&entity.Role{})
	if name != "" {
		db = db.Where("role_name LIKE ?", "%"+name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset((page - 1) * limit).Limit(limit).Find(&list).Error
	return list, total, err
}

// GetByID 单条
func (r *RoleRepo) GetByID(ctx context.Context, id uint64) (*entity.Role, error) {
	var role entity.Role
	err := DB.WithContext(ctx).First(&role, id).Error
	return &role, err
}

// Create 创建
func (r *RoleRepo) Create(ctx context.Context, role *entity.Role) error {
	return DB.WithContext(ctx).Create(role).Error
}

// Update 更新
func (r *RoleRepo) Update(ctx context.Context, role *entity.Role, id uint64) error {
	db := DB.WithContext(ctx).
		Model(&entity.Role{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(role)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}

// Delete 级联清权限
func (r *RoleRepo) Delete(ctx context.Context, id uint64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", id).Delete(&entity.RolePermissionRel{}).Error; err != nil {
			return err
		}
		if err := tx.Where("role_id = ?", id).Delete(&entity.GroupRoleRel{}).Error; err != nil {
			return err
		}
		return tx.Delete(&entity.Role{}, id).Error
	})
}

// ExistsCode 排除自身查重
func (r *RoleRepo) ExistsCode(ctx context.Context, code string, excludeID uint64) (bool, error) {
	var c int64
	err := DB.WithContext(ctx).Model(&entity.Role{}).
		Where("role_code = ? AND id <> ?", code, excludeID).Count(&c).Error
	return c > 0, err
}
