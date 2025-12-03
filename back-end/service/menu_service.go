/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:46:55
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 19:05:31
 * @FilePath: \back-end\service\menu_service.go
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
	"ginManager/utils"
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
	return s.repo.Update(ctx, &m, id)
}

// Delete 删除
func (s *MenuService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// GetByID 单条
func (s *MenuService) GetByID(ctx context.Context, id uint64) (*entity.Menu, error) {
	return s.repo.GetByID(ctx, id)
}

// service/menu_service.go
func (s *MenuService) SideMenu(ctx context.Context) (*dto.SideMenuResp, error) {
	// 1. 取启用菜单，按 sort 升序
	list, err := s.repo.Tree(ctx) // 已按 ParentID 升序
	if err != nil {
		return nil, err
	}

	// 2. 拼装 LayuiAdmin 格式
	root := &dto.SideMenuResp{
		HomeInfo: dto.HomeInfo{Title: "首页", Href: "page/welcome-1.html?t=1"},
		LogoInfo: dto.LogoInfo{Title: "LAYUI MINI", Image: "images/logo.png", Href: ""},
		MenuInfo: buildSideMenu(list),
	}
	return root, nil
}

// 把 entity.Menu → MenuNode 递归树
func buildSideMenu(list []entity.Menu) []dto.MenuNode {
	m := make(map[uint64][]dto.MenuNode)
	for _, v := range list {
		node := dto.MenuNode{
			ID:     v.ID,
			Title:  v.Title,
			Icon:   utils.StringPtrVal(v.Icon),
			Href:   *v.Path,
			Target: "_self", // 固定
		}
		m[v.ParentID] = append(m[v.ParentID], node)
	}
	var dfs func(parentID uint64) []dto.MenuNode
	dfs = func(parentID uint64) []dto.MenuNode {
		var res []dto.MenuNode
		for _, v := range m[parentID] {
			v.Child = dfs(v.ID) // 子级
			res = append(res, v)
		}
		return res
	}
	return dfs(0)
}

// service/perm_service.go
func BuildPermTree(list []entity.Menu, checkedMap map[uint64]bool) []dto.PermNode {
	// ① 分组：统一 uint64
	m := make(map[uint64][]dto.PermNode)
	for _, v := range list {
		node := dto.PermNode{
			AuthorityId:   int64(v.ID), // 保持 uint64
			AuthorityName: v.Title,
			OrderNumber:   v.Sort,
			MenuUrl:       v.Path,
			MenuIcon:      utils.StringPtrVal(v.Icon),
			CreateTime:    v.CreatedAt.Format("2006/01/02 15:04:05"),
			UpdateTime:    v.UpdatedAt.Format("2006/01/02 15:04:05"),
			ParentId:      int64(v.ParentID), // uint64
			IsMenu:        0,
			Checked:       0,
		}
		if checkedMap[v.ID] {
			node.Checked = 1
		}
		m[v.ParentID] = append(m[v.ParentID], node)
	}

	// ② 递归：统一 uint64
	var dfs func(parentID uint64) []dto.PermNode
	dfs = func(parentID uint64) []dto.PermNode {
		var res []dto.PermNode
		for _, v := range m[parentID] {
			v.Authority = dfs(uint64(v.AuthorityId)) // AuthorityId 也是 uint64
			res = append(res, v)
		}
		return res
	}

	// ③ 根节点：ParentID = 0（不是 -1）
	return dfs(0)
}
