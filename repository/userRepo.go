/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:37:48
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 13:50:35
 * @FilePath: \ginManager\repository\userRepo.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package repository

import (
	"context"
	"errors"
	"ginManager/models/entity"

	"gorm.io/gorm"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo { return &UserRepo{} }

// GetByUsername 根据用户名查用户
func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var u entity.User
	err := DB.WithContext(ctx).Where("username = ?", username).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

// Create 创建用户
func (r *UserRepo) Create(ctx context.Context, user *entity.User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// GetPage 分页+模糊查询
func (r *UserRepo) GetPage(ctx context.Context, username string, status uint8, page, limit int) (list []entity.User, total int64, err error) {
	db := DB.WithContext(ctx).Model(&entity.User{})
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}
	if status < 2 { // 0/1
		db = db.Where("status = ?", status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset((page - 1) * limit).Limit(limit).Find(&list).Error
	return list, total, err
}
