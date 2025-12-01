/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 09:23:18
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 09:27:50
 * @FilePath: \ginManager\service\permission_service.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package service

import (
	"context"
	"ginManager/models/entity"
	"ginManager/repository"
)

type PermissionService struct {
	repo *repository.RolePermRepo
}

func NewPermissionService() *PermissionService {
	return &PermissionService{repo: repository.NewRolePermRepo()}
}

// List 全部可用权限
func (s *PermissionService) List(ctx context.Context) ([]entity.Permission, error) {
	return s.repo.List(ctx)
}

// GetByIDs 批量
func (s *PermissionService) GetByIDs(ctx context.Context, ids []uint64) ([]entity.Permission, error) {
	return s.repo.GetByIDs(ctx, ids)
}
