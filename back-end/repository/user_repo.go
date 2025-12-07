/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:37:48
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-06 19:53:45
 * @FilePath: \back-end\repository\user_repo.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package repository

import (
	"context"
	"errors"
	"fmt"
	"ginManager/models/entity"
	"time"

	"gorm.io/gorm"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo { return &UserRepo{} }

// GetByID 根据主键查用户
func (r *UserRepo) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	var u entity.User
	err := DB.WithContext(ctx).Select("id", "username", "nickname", "email", "status", "created_at", "updated_at").First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

// GetByUsername 根据用户名查用户（已写）
func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var u entity.User
	err := DB.WithContext(ctx).Where("username = ?", username).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

// Page 分页+多条件查询（username/nickname/status/更新时间/删除状态）
func (r *UserRepo) Page(ctx context.Context,
	username, nickname string,
	status uint8,
	updatedStart, updatedEnd *time.Time,
	deleted *uint8, // nil 表示不筛选；0 未删 1 已删
	page, limit int) (list []entity.User, total int64, err error) {
	db := DB.WithContext(ctx).
		Model(&entity.User{}).
		Unscoped().
		Select("id", "username", "nickname", "email", "status", "created_at", "updated_at", "deleted_at")

	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}
	if nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	fmt.Println(status)
	if status < 2 { // 0/1
		db = db.Where("status = ?", status)
	}
	if updatedStart != nil && !updatedStart.IsZero() {
		db = db.Where("updated_at >= ?", *updatedStart)
	}
	if updatedEnd != nil && !updatedEnd.IsZero() {
		db = db.Where("updated_at <= ?", *updatedEnd)
	}
	if deleted != nil {
		if *deleted == 1 {
			db = db.Where("deleted_at IS NOT NULL")
		} else {
			db = db.Where("deleted_at IS NULL")
		}
	}

	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = db.Offset((page - 1) * limit).Limit(limit).Find(&list).Error
	return list, total, err
}

// Create 创建用户
func (r *UserRepo) Create(ctx context.Context, u *entity.User) error {
	return DB.WithContext(ctx).Create(u).Error
}

// Update 更新非零字段
func (r *UserRepo) Update(ctx context.Context, u *entity.User, id uint64) error {
	// 如果是管理员用户，强制保持激活状态
	var old entity.User
	if err := DB.WithContext(ctx).First(&old, id).Error; err != nil {
		return err
	}
	if old.Username == "admin" {
		u.Status = 1
	}

	db := DB.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(u)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}

// UpdateStatus 单独切换状态
func (r *UserRepo) UpdateStatus(ctx context.Context, id uint64, status uint8) error {
	// 禁止修改管理员用户的状态
	var u entity.User
	if err := DB.WithContext(ctx).Select("id", "username").First(&u, id).Error; err != nil {
		return err
	}
	if u.Username == "admin" {
		return errors.New("管理员用户状态不能被更改")
	}

	db := DB.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("status", status)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}

// UpdatePassword 单独改密码
func (r *UserRepo) UpdatePassword(ctx context.Context, id uint64, newHash string) error {
	db := DB.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("password_hash", newHash)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}

// Delete 物理删除（可选逻辑删）
func (r *UserRepo) Delete(ctx context.Context, id uint64) error {
	return DB.WithContext(ctx).
		Where("username != ?", "admin").
		Delete(&entity.User{}, id).Error
}

// ExistsUsername 排除自身查重
func (r *UserRepo) ExistsUsername(ctx context.Context, username string, excludeID uint64) (bool, error) {
	var c int64
	err := DB.WithContext(ctx).Model(&entity.User{}).
		Where("username = ? AND id <> ?", username, excludeID).Count(&c).Error
	return c > 0, err
}
