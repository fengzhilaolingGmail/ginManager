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
	return s.repo.Page(ctx, req.GroupName, req.Page, req.Limit)
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
		GroupCode: req.GroupCode,
		GroupName: req.GroupName,
		Sort:      req.Sort,
		Status:    req.Status,
	}
	return s.repo.Create(ctx, &g)
}

// Update 编辑
func (s *GroupService) Update(ctx context.Context, id uint64, req *dto.GroupAddReq) error {
	exist, err := s.repo.ExistsCode(ctx, req.GroupCode, id)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("组编码已存在")
	}
	g := entity.UserGroup{
		ID:        id,
		GroupCode: req.GroupCode,
		GroupName: req.GroupName,
		Sort:      req.Sort,
		Status:    req.Status,
	}
	return s.repo.Update(ctx, &g, id)
}

// Delete 删除
func (s *GroupService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// GetByID 单条
func (s *GroupService) GetByID(ctx context.Context, id uint64) (*entity.UserGroup, error) {
	return s.repo.GetByID(ctx, id)
}
