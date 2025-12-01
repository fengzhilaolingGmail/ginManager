/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:48:14
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 16:55:26
 * @FilePath: \ginManager\dto\request.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package dto

// =========== 登录 ===========
type LoginReq struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// =========== 通用分页 ===========
type PageReq struct {
	Page  int `form:"page"  binding:"required,min=1"` // 第几页
	Limit int `form:"limit" binding:"required,min=1,max=200"`
}

// =========== 用户管理 ===========
type UserAddReq struct {
	Username string  `json:"username" binding:"required,min=3,max=20"`
	Password string  `json:"password" binding:"required,min=6,max=32"`
	Nickname string  `json:"nickname" binding:"required,min=1,max=50"`
	Email    *string `json:"email"    binding:"omitempty,email,max=100"`
	Phone    *string `json:"phone"    binding:"omitempty,max=20"`
	Status   uint8   `json:"status"   binding:"oneof=0 1"`
}

type UserUpdateReq struct {
	Username string  `json:"username" binding:"omitempty,min=3,max=32"`
	Nickname string  `json:"nickname" binding:"omitempty,max=32"`
	Email    *string `json:"email"    binding:"omitempty,email"`
	Status   *uint8  `json:"status"   binding:"omitempty,oneof=0 1"` // 指针才能区分“没传”和“传0”
	Phone    *string `json:"phone"    binding:"omitempty,max=20"`
	Password string  `json:"password" binding:"omitempty,min=6,max=64"` // 修改密码时必填
}

type UserListReq struct {
	PageReq
	Username string `form:"username"` // 模糊查询
	Status   uint8  `form:"status"`   // 0/1/ 空=全部
}
