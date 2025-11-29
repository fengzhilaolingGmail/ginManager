/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:37:07
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 13:37:20
 * @FilePath: \ginManager\repository\repo.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package repository

import "gorm.io/gorm"

var DB *gorm.DB

// SetDB 由 main 注入
func SetDB(db *gorm.DB) {
	DB = db
}
