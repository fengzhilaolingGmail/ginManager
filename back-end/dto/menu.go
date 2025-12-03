/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:42:30
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 15:20:08
 * @FilePath: \back-end\dto\menu.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// dto/menu.go
package dto

type MenuAddReq struct {
	ParentID  uint64  `json:"parent_id"`
	Title     string  `json:"title" binding:"required,max=50"`
	Name      *string `json:"name"`      // 路由 name
	Path      *string `json:"path"`      // 前端路由
	Component *string `json:"component"` // 前端组件
	Sort      int     `json:"sort"`
	Status    uint8   `json:"status" binding:"oneof=0 1"`
}

type SideMenuResp struct {
	HomeInfo HomeInfo   `json:"homeInfo"`
	LogoInfo LogoInfo   `json:"logoInfo"`
	MenuInfo []MenuNode `json:"menuInfo"`
}

type HomeInfo struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

type LogoInfo struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Href  string `json:"href"`
}

type MenuNode struct {
	ID     uint64     `json:"-"` // 不返回前端
	Title  string     `json:"title"`
	Icon   string     `json:"icon"`
	Href   string     `json:"href"`
	Target string     `json:"target"`
	Child  []MenuNode `json:"child,omitempty"` // 子级
}

type PermNode struct {
	AuthorityId   int64      `json:"authorityId"`   // 主键
	AuthorityName string     `json:"authorityName"` // 标题
	OrderNumber   int        `json:"orderNumber"`   // 排序
	MenuUrl       *string    `json:"menuUrl"`       // 前端路由
	MenuIcon      string     `json:"menuIcon"`      // 图标
	CreateTime    string     `json:"createTime"`    // 创建时间
	UpdateTime    string     `json:"updateTime"`    // 更新时间
	ParentId      int64      `json:"parentId"`      // 父节点
	IsMenu        int        `json:"isMenu"`        // 0=目录 1=菜单 2=按钮
	Checked       int        `json:"checked"`       // 是否选中（角色已拥有）
	Authority     []PermNode `json:"authority"`     // 子级
}
