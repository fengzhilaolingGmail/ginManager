/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 09:22:22
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 15:45:45
 * @FilePath: \back-end\service\role_service.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package service

import (
	"context"
	"errors"
	"ginManager/dto"
	"ginManager/models/entity"
	"ginManager/repository"
)

type RoleService struct {
	repo     *repository.RoleRepo
	permRepo *repository.RolePermRepo
	rolePerm *repository.RolePermRepo
}

func NewRoleService() *RoleService {
	return &RoleService{
		repo:     repository.NewRoleRepo(),
		permRepo: repository.NewRolePermRepo(),
		rolePerm: repository.NewRolePermRepo(),
	}
}

// Page 分页
func (s *RoleService) Page(ctx context.Context, req *dto.RoleListReq) ([]entity.Role, int64, error) {
	return s.repo.Page(ctx, req.RoleName, req.Page, req.Limit)
}

// Create 新增角色 + 绑定权限
func (s *RoleService) Create(ctx context.Context, req *dto.RoleAddReq) error {
	exist, err := s.repo.ExistsCode(ctx, req.RoleCode, 0)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("角色编码已存在")
	}
	role := entity.Role{
		RoleCode: req.RoleCode,
		RoleName: req.RoleName,
		Sort:     req.Sort,
		Status:   req.Status,
	}
	if err := s.repo.Create(ctx, &role); err != nil {
		return err
	}
	// 绑定权限
	return s.rolePerm.SetPermissions(ctx, role.ID, req.PermIDs)
}

// Update 编辑角色 + 重绑权限
func (s *RoleService) Update(ctx context.Context, id uint64, req *dto.RoleAddReq) error {
	exist, err := s.repo.ExistsCode(ctx, req.RoleCode, id)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("角色编码已存在")
	}
	role := entity.Role{
		ID:       id,
		RoleCode: req.RoleCode,
		RoleName: req.RoleName,
		Sort:     req.Sort,
		Status:   req.Status,
	}
	if err := s.repo.Update(ctx, &role, id); err != nil {
		return err
	}
	return s.rolePerm.SetPermissions(ctx, id, req.PermIDs)
}

// Delete 删除角色
func (s *RoleService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// Get 单条 + 已有权限 ID 数组
func (s *RoleService) Get(ctx context.Context, id uint64) (*entity.Role, []uint64, error) {
	role, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	permIDs, err := s.rolePerm.GetPermIDsByRole(ctx, id)
	return role, permIDs, err
}
