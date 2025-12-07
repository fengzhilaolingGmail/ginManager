package service

import (
	"context"
	"errors"
	"fmt"
	"ginManager/dto"
	"ginManager/models/entity"
	"ginManager/repository"
	"ginManager/utils"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService() *UserService {
	return &UserService{repo: repository.NewUserRepo()}
}

// Page 分页列表
func (s *UserService) Page(ctx context.Context, req *dto.UserListReq) ([]entity.User, int64, error) {
	return s.repo.Page(ctx,
		req.Username,
		req.Nickname,
		req.Status,
		req.UpdatedStart,
		req.UpdatedEnd,
		req.Deleted,
		req.Page,
		req.Limit)
}

// Create 新增用户
func (s *UserService) Create(ctx context.Context, req *dto.UserAddReq) error {
	// 查重
	exist, err := s.repo.ExistsUsername(ctx, req.Username, 0)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户名已存在")
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	u := entity.User{
		Username:     req.Username,
		PasswordHash: hash,
		Nickname:     req.Nickname,
		Email:        req.Email,
		Phone:        req.Phone,
		Status:       req.Status,
	}
	return s.repo.Create(ctx, &u)
}

// Update 编辑用户（不改密码）
func (s *UserService) Update(ctx context.Context, id uint64, req *dto.UserUpdateReq) error {
	exist, err := s.repo.ExistsUsername(ctx, req.Username, id)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户名已存在")
	}
	// 构造要更新的用户实体（不包含 status），status 单独处理以支持设置为 0
	u := entity.User{
		ID:       id,
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	// 先更新普通字段（status 不在此处）
	if err := s.repo.Update(ctx, &u, id); err != nil {
		return err
	}
	fmt.Println(req.Status)
	// 如果请求中包含 status，则单独调用 UpdateStatus（支持设置为 0）
	if req.Status != nil {
		if err := s.repo.UpdateStatus(ctx, id, *req.Status); err != nil {
			return err
		}
	}
	return nil
}

// UpdateStatus 开关账号
func (s *UserService) UpdateStatus(ctx context.Context, id uint64, status uint8) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

// UpdatePassword 修改密码
func (s *UserService) UpdatePassword(ctx context.Context, id uint64, newPwd string) error {
	hash, err := utils.HashPassword(newPwd)
	if err != nil {
		return err
	}
	return s.repo.UpdatePassword(ctx, id, hash)
}

// Delete 删除用户
func (s *UserService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// GetByID 单条详情
func (s *UserService) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	return s.repo.GetByID(ctx, id)
}
