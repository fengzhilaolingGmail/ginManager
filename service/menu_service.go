/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:46:55
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-30 10:03:00
 * @FilePath: \ginManager\service\menu_service.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// service/menu_service.go
package service

import (
	"context"
	"ginManager/dto"
	"ginManager/models/entity"
	"ginManager/repository"
)

type MenuService struct {
	repo *repository.MenuRepo
}

func NewMenuService() *MenuService {
	return &MenuService{repo: repository.NewMenuRepo()}
}

// Tree 树形菜单
func (s *MenuService) Tree(ctx context.Context) ([]entity.Menu, error) {
	list, err := s.repo.Tree(ctx)
	if err != nil {
		return nil, err
	}
	// 组装树
	return buildTree(list), nil
}

// 内部递归
// buildTree 把扁平菜单列表转成树（parentID = 0 为根）
func buildTree(list []entity.Menu) []entity.Menu {
	// 1. 按 parentID 分组
	m := make(map[uint64][]entity.Menu)
	for _, v := range list {
		m[v.ParentID] = append(m[v.ParentID], v)
	}

	// 2. 从根开始递归
	var dfs func(parentID uint64) []entity.Menu
	dfs = func(parentID uint64) []entity.Menu {
		var res []entity.Menu
		for _, v := range m[parentID] {
			v.Children = dfs(v.ID) // 组装子树
			res = append(res, v)
		}
		return res
	}
	return dfs(0) // 入口：根节点 parentID = 0
}

// Create 新增
func (s *MenuService) Create(ctx context.Context, req *dto.MenuAddReq) error {
	m := entity.Menu{
		ParentID:  req.ParentID,
		Title:     req.Title,
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Sort:      req.Sort,
		Status:    req.Status,
	}
	return s.repo.Create(ctx, &m)
}

// Update 编辑
func (s *MenuService) Update(ctx context.Context, id uint64, req *dto.MenuAddReq) error {
	m := entity.Menu{
		ID:        id,
		ParentID:  req.ParentID,
		Title:     req.Title,
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Sort:      req.Sort,
		Status:    req.Status,
	}
	return s.repo.Update(ctx, &m)
}

// Delete 删除
func (s *MenuService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// GetByID 单条
func (s *MenuService) GetByID(ctx context.Context, id uint64) (*entity.Menu, error) {
	return s.repo.GetByID(ctx, id)
}
