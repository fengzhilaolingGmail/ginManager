// repository/oper_log_repo.go
package repository

import (
	"context"
	"fmt"
	"ginManager/dto"
	"ginManager/models/entity"
)

type UserLogRepo struct{}

func NewUserLogRepo() *UserLogRepo { return &UserLogRepo{} }

// Create 插入单条
func (r *UserLogRepo) Create(ctx context.Context, log *entity.UserLog) error {
	return DB.WithContext(ctx).Create(log).Error
}

// Page 分页 + 模糊查询
func (r *UserLogRepo) Page(ctx context.Context, req *dto.UserLogListReq) (list []entity.UserLog, total int64, err error) {
	db := DB.WithContext(ctx).Model(&entity.UserLog{})
	if req.Module != "" {
		db = db.Where("module LIKE ?", "%"+req.Module+"%")
		fmt.Println("module:", req.Module)
	}
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Status < 2 { // 0/1
		db = db.Where("status = ?", req.Status)
	}
	if !req.StartTime.IsZero() && !req.EndTime.IsZero() {
		db = db.Where("created_at BETWEEN ? AND ?", req.StartTime, req.EndTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("created_at DESC").Offset((req.Page - 1) * req.Limit).Limit(req.Limit).Find(&list).Error
	return list, total, err
}
