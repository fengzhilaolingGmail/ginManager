/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:44:24
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 10:37:18
 * @FilePath: \ginManager\repository\menu_repo.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// repository/menu_repo.go
package repository

import (
	"context"
	"ginManager/models/entity"

	"gorm.io/gorm"
)

type MenuRepo struct{}

func NewMenuRepo() *MenuRepo { return &MenuRepo{} }

// Tree 返回树形（parent_id=0 为根）
func (r *MenuRepo) Tree(ctx context.Context) ([]entity.Menu, error) {
	var list []entity.Menu
	err := DB.WithContext(ctx).Where("status = 1").Order("sort asc").Find(&list).Error
	return list, err
}

// GetByID 单条
func (r *MenuRepo) GetByID(ctx context.Context, id uint64) (*entity.Menu, error) {
	var m entity.Menu
	err := DB.WithContext(ctx).First(&m, id).Error
	return &m, err
}

// Create 创建
func (r *MenuRepo) Create(ctx context.Context, m *entity.Menu) error {
	return DB.WithContext(ctx).Create(m).Error
}

// Update 更新
func (r *MenuRepo) Update(ctx context.Context, m *entity.Menu, id uint64) error {
	db := DB.WithContext(ctx).
		Model(&entity.Menu{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(m)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}

// Delete 删除（含子节点）
func (r *MenuRepo) Delete(ctx context.Context, id uint64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先删子节点
		if err := tx.Where("parent_id = ?", id).Delete(&entity.Menu{}).Error; err != nil {
			return err
		}
		// 再删自己
		return tx.Delete(&entity.Menu{}, id).Error
	})
}
