package repository

import (
	"context"
	"ginManager/models/entity"

	"gorm.io/gorm"
)

type RolePermRepo struct{}

func NewRolePermRepo() *RolePermRepo { return &RolePermRepo{} }

// SetPermissions 给角色重新绑定权限（先删后插）
func (r *RolePermRepo) SetPermissions(ctx context.Context, roleID uint64, permIDs []uint64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleID).Delete(&entity.RolePermissionRel{}).Error; err != nil {
			return err
		}
		if len(permIDs) == 0 {
			return nil
		}
		rels := make([]entity.RolePermissionRel, len(permIDs))
		for i, pid := range permIDs {
			rels[i] = entity.RolePermissionRel{RoleID: roleID, PermID: pid}
		}
		return tx.CreateInBatches(rels, 200).Error
	})
}

// GetPermIDsByRole 查询角色已有权限 ID 数组
func (r *RolePermRepo) GetPermIDsByRole(ctx context.Context, roleID uint64) ([]uint64, error) {
	var ids []uint64
	err := DB.WithContext(ctx).
		Model(&entity.RolePermissionRel{}).
		Where("role_id = ?", roleID).
		Pluck("perm_id", &ids).Error
	return ids, err
}

// List 查询所有启用的权限，按 sort 升序
func (r *RolePermRepo) List(ctx context.Context) ([]entity.Permission, error) {
	var list []entity.Permission
	err := DB.WithContext(ctx).
		Where("status = 1").
		Order("sort ASC").
		Find(&list).Error
	return list, err
}

// GetByIDs 根据 ID 批量查询权限
func (r *RolePermRepo) GetByIDs(ctx context.Context, ids []uint64) ([]entity.Permission, error) {
	var list []entity.Permission
	err := DB.WithContext(ctx).
		Where("id IN ?", ids).
		Find(&list).Error
	return list, err
}
