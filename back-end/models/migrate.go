/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:04:52
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 12:11:07
 * @FilePath: \ginManager\models\migrate.go
 * @Description: db 连接与自动迁移
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package models

import (
	"ginManager/config"
	"ginManager/models/entity"

	"ginManager/logger"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gromLogger "gorm.io/gorm/logger"
)

// var DB *gorm.DB

// Init 初始化连接并自动迁移
func Init() *gorm.DB {
	var dia gorm.Dialector
	switch config.C.DB.Driver {
	case "mysql":
		dia = mysql.Open(config.C.DB.DSN)
	case "postgres":
		dia = postgres.Open(config.C.DB.DSN)
	default:
		logger.L.Fatal("unsupported driver", zap.String("driver", config.C.DB.Driver))
	}

	var err error
	db, err := gorm.Open(dia, &gorm.Config{
		Logger: logger.NewGormLogger(logger.L).LogMode(gromLogger.Info), // 把 zap 接入 gorm
	})
	if err != nil {
		logger.L.Fatal("connect db fail", zap.Error(err))
	}

	if err = autoMigrate(db); err != nil {
		logger.L.Fatal("auto migrate fail", zap.Error(err))
	}
	logger.L.Info("db connected and migrated")
	// 首次初始化数据
	InitFirstRun(db)
	InitMenuPerm(db)
	return db
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.UserGroup{},
		&entity.UserGroupRel{},
		&entity.Role{},
		&entity.GroupRoleRel{},
		&entity.Permission{},
		&entity.RolePermissionRel{},
		&entity.Menu{},
		&entity.PermissionMenuRel{},
		&entity.UserLog{},
	)
}
