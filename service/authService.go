/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:46:25
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 13:52:23
 * @FilePath: \ginManager\service\authService.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// service/auth_service.go
package service

import (
	"context"
	"errors"
	"ginManager/dto"
	"ginManager/models/entity"
	"ginManager/repository"
	"ginManager/utils"
)

type AuthService struct {
	userRepo *repository.UserRepo
	authRepo *repository.AuthRepo
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepo(),
		authRepo: repository.NewAuthRepo(),
	}
}

// Login 登录并返回 token
func (s *AuthService) Login(ctx context.Context, req *dto.LoginReq) (token string, user *entity.User, err error) {
	u, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return "", nil, err
	}
	if u == nil || u.Status == 0 {
		return "", nil, errors.New("账号或密码错误")
	}

	if !utils.CheckPassword(req.Password, u.PasswordHash) {
		return "", nil, errors.New("账号或密码错误")
	}
	token, err = utils.GenerateToken(u.ID)
	return token, u, err
}

// GetPermissions 获取用户权限列表
func (s *AuthService) GetPermissions(ctx context.Context, userID uint64) ([]string, error) {
	return s.authRepo.GetPermissionCodes(ctx, userID)
}
