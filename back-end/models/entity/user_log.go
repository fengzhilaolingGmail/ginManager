/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:10:04
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 11:02:02
 * @FilePath: \ginManager\models\entity\user_log.go
 * @Description: 用户操作日志表
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package entity

import (
	"time"

	"gorm.io/gorm"
)

// ----------------- 用户操作日志 -----------------
type UserLog struct {
	ID         uint64  `gorm:"primaryKey;autoIncrement;comment:日志主键"`
	UserID     *uint64 `gorm:"comment:用户ID"`
	Username   *string `gorm:"size:50;comment:用户名"`
	Module     *string `gorm:"size:50;comment:模块"`
	Action     *string `gorm:"size:50;comment:动作"`
	Method     *string `gorm:"size:10;comment:HTTP方法"`
	Path       *string `gorm:"size:200;comment:请求路径"`
	IP         *string `gorm:"size:45;comment:客户端IP"`
	UserAgent  *string `gorm:"type:text;comment:UA"`
	Status     uint8   `gorm:"not null;default:1;comment:1成功 0失败"`
	ErrorMsg   *string `gorm:"type:text;comment:错误信息"`
	DurationMs *int    `gorm:"comment:耗时(ms)"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"` // 软删除
}
