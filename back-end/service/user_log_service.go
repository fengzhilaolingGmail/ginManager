/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 10:54:52
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 12:19:45
 * @FilePath: \ginManager\service\user_log_service.go
 * @Description: 操作日志记录服务
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package service

import (
	"context"
	"ginManager/dto"
	"ginManager/models/entity"
	"ginManager/repository"
)

type UserLogService struct {
	repo *repository.UserLogRepo
}

func NewUserLogService() *UserLogService {
	return &UserLogService{repo: repository.NewUserLogRepo()}
}

// Page 分页列表
func (s *UserLogService) Page(ctx context.Context, req *dto.UserLogListReq) ([]entity.UserLog, int64, error) {
	return s.repo.Page(ctx, req)
}

// Create 内部调用（中间件用）
func (s *UserLogService) Create(ctx context.Context, log *entity.UserLog) error {
	return s.repo.Create(ctx, log)
}
