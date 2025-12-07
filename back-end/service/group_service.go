/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:46:08
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 10:41:02
 * @FilePath: \ginManager\service\group_service.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// service/group_service.go
package service

import (
	"context"
	"errors"
	"ginManager/dto"
	"ginManager/models/entity"
	"ginManager/repository"
)

type GroupService struct {
	repo *repository.GroupRepo
}

func NewGroupService() *GroupService {
	return &GroupService{repo: repository.NewGroupRepo()}
}

// Page 分页
func (s *GroupService) Page(ctx context.Context, req *dto.GroupListReq) ([]entity.UserGroup, int64, error) {
	// UpdatedStart/UpdatedEnd are already *time.Time in DTO, pass through
	return s.repo.Page(ctx, req.GroupName, req.GroupCode, req.Status, req.UpdatedStart, req.UpdatedEnd, req.Deleted, req.Page, req.Limit)
}

// Create 新增
func (s *GroupService) Create(ctx context.Context, req *dto.GroupAddReq) error {
	exist, err := s.repo.ExistsCode(ctx, req.GroupCode, 0)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("组编码已存在")
	}
	g := entity.UserGroup{
		GroupCode:   req.GroupCode,
		GroupName:   req.GroupName,
		Sort:        req.Sort,
		Status:      req.Status,
		Description: req.Description,
	}
	return s.repo.Create(ctx, &g)
}

// Update 编辑
func (s *GroupService) Update(ctx context.Context, id uint64, req *dto.GroupUpdateReq) error {
	exist, err := s.repo.ExistsCode(ctx, req.GroupCode, id)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("组编码已存在")
	}

	// 构造更新实体（不包含 status），使用单独的 UpdateStatus 来处理可能的零值
	g := entity.UserGroup{
		ID:        id,
		GroupCode: req.GroupCode,
		GroupName: req.GroupName,
	}
	if req.Sort != nil {
		g.Sort = *req.Sort
	}
	if req.Description != nil {
		g.Description = *req.Description
	}

	if err := s.repo.Update(ctx, &g, id); err != nil {
		return err
	}

	if req.Status != nil {
		if err := s.repo.UpdateStatus(ctx, id, *req.Status); err != nil {
			return err
		}
	}
	return nil
}

// Delete 删除
func (s *GroupService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// UpdateStatus 切换状态
func (s *GroupService) UpdateStatus(ctx context.Context, id uint64, status uint8) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

// GetByID 单条
func (s *GroupService) GetByID(ctx context.Context, id uint64) (*entity.UserGroup, error) {
	return s.repo.GetByID(ctx, id)
}

// GetRolesPermsByGroup 查询组下角色和角色对应权限并转换为前端树形结构
func (s *GroupService) GetRolesPermsByGroup(ctx context.Context, groupID uint64) ([]dto.RoleItem, error) {
	roleRepo := repository.NewRoleRepo()
	items, err := roleRepo.GetRolesWithPermsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	var out []dto.RoleItem
	for _, it := range items {
		ri := dto.RoleItem{
			ID:       it.Role.ID,
			RoleCode: it.Role.RoleCode,
			RoleName: it.Role.RoleName,
			Status:   it.Role.Status,
		}
		for _, p := range it.Permissions {
			ri.Children = append(ri.Children, dto.PermItem{ID: p.ID, PermCode: p.PermCode, PermName: p.PermName})
		}
		out = append(out, ri)
	}
	return out, nil
}
