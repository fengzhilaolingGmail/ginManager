/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:38:29
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 13:38:42
 * @FilePath: \ginManager\repository\auth_repo.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package repository

import (
	"context"
	"ginManager/models/entity"
)

type AuthRepo struct{}

func NewAuthRepo() *AuthRepo { return &AuthRepo{} }

// GetPermissionCodes 获取用户所有权限编码
func (r *AuthRepo) GetPermissionCodes(ctx context.Context, userID uint64) ([]string, error) {
	var codes []string
	err := DB.WithContext(ctx).
		Model(&entity.Permission{}).
		Joins("JOIN role_permission_rel ON role_permission_rel.perm_id = permission.id").
		Joins("JOIN group_role_rel ON group_role_rel.role_id = role_permission_rel.role_id").
		Joins("JOIN user_group_rel ON user_group_rel.group_id = group_role_rel.group_id").
		Where("user_group_rel.user_id = ?", userID).
		Pluck("DISTINCT permission.perm_code", &codes).Error
	return codes, err
}
