/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:43:36
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 10:40:49
 * @FilePath: \ginManager\repository\group_repo.go
 * @Description: 用户组仓库
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// repository/group_repo.go
package repository

import (
	"context"
	"ginManager/models/entity"

	"gorm.io/gorm"
)

type GroupRepo struct{}

func NewGroupRepo() *GroupRepo { return &GroupRepo{} }

// Page 分页
func (r *GroupRepo) Page(ctx context.Context, name string, page, limit int) (list []entity.UserGroup, total int64, err error) {
	db := DB.WithContext(ctx).Model(&entity.UserGroup{})
	if name != "" {
		db = db.Where("group_name LIKE ?", "%"+name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset((page - 1) * limit).Limit(limit).Find(&list).Error
	return list, total, err
}

// GetByID 单条
func (r *GroupRepo) GetByID(ctx context.Context, id uint64) (*entity.UserGroup, error) {
	var g entity.UserGroup
	err := DB.WithContext(ctx).First(&g, id).Error
	return &g, err
}

// Create 创建
func (r *GroupRepo) Create(ctx context.Context, g *entity.UserGroup) error {
	return DB.WithContext(ctx).Create(g).Error
}

// Update 更新
func (r *GroupRepo) Update(ctx context.Context, g *entity.UserGroup, id uint64) error {
	db := DB.WithContext(ctx).
		Model(&entity.UserGroup{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(g)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}

// Delete 删除（同步删除关联表）
func (r *GroupRepo) Delete(ctx context.Context, id uint64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("group_id = ?", id).Delete(&entity.UserGroupRel{}).Error; err != nil {
			return err
		}
		if err := tx.Where("group_id = ?", id).Delete(&entity.GroupRoleRel{}).Error; err != nil {
			return err
		}
		return tx.Delete(&entity.UserGroup{}, id).Error
	})
}

// ExistsCode 排除自身查重
func (r *GroupRepo) ExistsCode(ctx context.Context, code string, excludeID uint64) (bool, error) {
	var c int64
	err := DB.WithContext(ctx).Model(&entity.UserGroup{}).
		Where("group_code = ? AND id <> ?", code, excludeID).Count(&c).Error
	return c > 0, err
}
