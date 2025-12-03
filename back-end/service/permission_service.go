/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 09:23:18
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 18:45:53
 * @FilePath: \back-end\service\permission_service.go
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
	repo     *repository.RolePermRepo
	rolePerm *repository.RolePermRepo
}

func NewPermissionService() *PermissionService {
	return &PermissionService{
		repo:     repository.NewRolePermRepo(),
		rolePerm: repository.NewRolePermRepo(),
	}
}

// List 全部可用权限
func (s *PermissionService) List(ctx context.Context) ([]entity.Permission, error) {
	return s.repo.List(ctx)
}

// GetByIDs 批量
func (s *PermissionService) GetByIDs(ctx context.Context, ids []uint64) ([]entity.Permission, error) {
	return s.repo.GetByIDs(ctx, ids)
}

// GetPermIDsByRole 根据角色ID获取权限ID列表
func (s *PermissionService) GetPermIDsByRole(ctx context.Context, roleID uint64) ([]uint64, error) {
	return s.rolePerm.GetPermIDsByRole(ctx, roleID)
}
