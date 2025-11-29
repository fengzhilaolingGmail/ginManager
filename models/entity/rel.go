/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:10:14
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 09:13:13
 * @FilePath: \ginManager\models\entity\rel.go
 * @Description: 关联关系表
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package entity

// ----------------- 3. 用户 <=> 用户组 -----------------
type UserGroupRel struct {
	ID      uint64    `gorm:"primaryKey;autoIncrement"`
	UserID  uint64    `gorm:"not null;uniqueIndex:idx_user_group_rel,priority:1"` // 联合唯一
	GroupID uint64    `gorm:"not null;uniqueIndex:idx_user_group_rel,priority:2"`
	User    User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Group   UserGroup `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}

// ----------------- 5. 用户组 <=> 角色 -----------------
type GroupRoleRel struct {
	ID      uint64    `gorm:"primaryKey;autoIncrement"`
	GroupID uint64    `gorm:"not null;uniqueIndex:idx_group_role_rel,priority:1"`
	RoleID  uint64    `gorm:"not null;uniqueIndex:idx_group_role_rel,priority:2"`
	Group   UserGroup `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Role    Role      `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}

// ----------------- 7. 角色 <=> 权限 -----------------
type RolePermissionRel struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement"`
	RoleID     uint64     `gorm:"not null;uniqueIndex:idx_role_perm_rel,priority:1"`
	PermID     uint64     `gorm:"not null;uniqueIndex:idx_role_perm_rel,priority:2"`
	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}

// ----------------- 9. 权限 <=> 菜单 -----------------
type PermissionMenuRel struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement"`
	PermID     uint64     `gorm:"not null;uniqueIndex:idx_perm_menu_rel,priority:1"`
	MenuID     uint64     `gorm:"not null;uniqueIndex:idx_perm_menu_rel,priority:2"`
	Permission Permission `gorm:"foreignKey:PermID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Menu       Menu       `gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}
