package models

import (
	"context"
	"fmt"
	"ginManager/logger"
	"ginManager/models/entity"
	"ginManager/utils"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitFirstRun 首次启动初始化
func InitFirstRun(db *gorm.DB) {
	var cnt int64
	if err := db.Model(&entity.User{}).Count(&cnt).Error; err != nil {
		logger.L.Fatal("check user count fail", zap.Error(err))
	}
	if cnt > 0 {
		logger.L.Info("db already seeded, skip init")
		return
	}

	logger.L.Info("start first-run init")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pwd, err := utils.GeneratePassword(12) // 生成随机密码
	if err != nil {
		logger.L.Fatal("generate password fail", zap.Error(err))
	}
	pwdHash, err := utils.HashPassword(pwd) // 加密密码
	if err != nil {
		logger.L.Fatal("hash password fail", zap.Error(err))
	}
	logger.L.Info("generated admin password", zap.String("password", pwd))
	if err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 创建超级管理员账号
		admin := entity.User{
			Username:     "admin",
			PasswordHash: pwdHash, // 你自己封装的 bcrypt 工具
			Nickname:     "超级管理员",
			Email:        stringPtr("admin@local.dev"),
			Phone:        nil,
			Status:       1,
		}
		if err := tx.Create(&admin).Error; err != nil {
			return err
		}

		// 2. 创建「系统管理组」
		grp := entity.UserGroup{
			GroupCode: "sysadmin",
			GroupName: "系统管理组",
			Sort:      0,
			Status:    1,
		}
		if err := tx.Create(&grp).Error; err != nil {
			return err
		}

		// 3. 创建超级角色
		role := entity.Role{
			RoleCode: "super-admin",
			RoleName: "超级管理员角色",
			Sort:     0,
			Status:   1,
		}
		if err := tx.Create(&role).Error; err != nil {
			return err
		}

		// 4. 用户 -> 组
		if err := tx.Create(&entity.UserGroupRel{UserID: admin.ID, GroupID: grp.ID}).Error; err != nil {
			return err
		}

		// 5. 组 -> 角色
		if err := tx.Create(&entity.GroupRoleRel{GroupID: grp.ID, RoleID: role.ID}).Error; err != nil {
			return err
		}
		path := fmt.Sprintf("/%s", "system")
		// 6. 内置全部菜单（示例只插一条根菜单，可按需扩充）
		menu := entity.Menu{
			ParentID:  0,
			Title:     "系统管理",
			Path:      &path,
			Name:      stringPtr("System"),
			Component: stringPtr("Layout"),
			Sort:      0,
			Status:    1,
		}
		if err := tx.Create(&menu).Error; err != nil {
			return err
		}

		// 7. 创建「*」权限节点
		perm := entity.Permission{
			PermCode: "*",
			PermName: "所有权限",
			Sort:     0,
			Status:   1,
		}
		if err := tx.Create(&perm).Error; err != nil {
			return err
		}

		// 8. 角色 -> 权限
		if err := tx.Create(&entity.RolePermissionRel{RoleID: role.ID, PermID: perm.ID}).Error; err != nil {
			return err
		}

		// 9. 权限 -> 菜单
		if err := tx.Create(&entity.PermissionMenuRel{PermID: perm.ID, MenuID: menu.ID}).Error; err != nil {
			return err
		}

		logger.L.Info("first-run init success",
			zap.Uint64("adminId", admin.ID),
			zap.Uint64("groupId", grp.ID),
			zap.Uint64("roleId", role.ID))
		return nil
	}); err != nil {
		logger.L.Fatal("first-run init fail", zap.Error(err))
	}
}

// InitMenuPerm 初始化菜单与权限数据
func InitMenuPerm(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 顶级菜单
		sysMgr := entity.Menu{ParentID: 0, Title: "系统管理", Name: stringPtr("System"), Path: stringPtr("/system"), Component: stringPtr("Layout"), Sort: 1, Status: 1}
		if err := tx.FirstOrCreate(&sysMgr, entity.Menu{Path: stringPtr("/system")}).Error; err != nil {
			return err
		}
		logMgr := entity.Menu{ParentID: 0, Title: "系统日志", Name: stringPtr("Log"), Path: stringPtr("/log"), Component: stringPtr("Layout"), Sort: 2, Status: 1}
		if err := tx.FirstOrCreate(&logMgr, entity.Menu{Path: stringPtr("/log")}).Error; err != nil {
			return err
		}

		// 2. 子菜单
		sub := []entity.Menu{
			{ParentID: sysMgr.ID, Title: "用户管理", Name: stringPtr("User"), Path: stringPtr("/system/user"), Component: stringPtr("system/user/index"), Sort: 11, Status: 1},
			{ParentID: sysMgr.ID, Title: "用户组管理", Name: stringPtr("UserGroup"), Path: stringPtr("/system/group"), Component: stringPtr("system/group/index"), Sort: 12, Status: 1},
			{ParentID: sysMgr.ID, Title: "菜单管理", Name: stringPtr("Menu"), Path: stringPtr("/system/menu"), Component: stringPtr("system/menu/index"), Sort: 13, Status: 1},
			{ParentID: sysMgr.ID, Title: "权限管理", Name: stringPtr("Perm"), Path: stringPtr("/system/perm"), Component: stringPtr("system/perm/index"), Sort: 14, Status: 1},
			{ParentID: logMgr.ID, Title: "操作日志", Name: stringPtr("OperLog"), Path: stringPtr("/log/oper"), Component: stringPtr("log/oper/index"), Sort: 21, Status: 1},
		}
		for i := range sub {
			if err := tx.FirstOrCreate(&sub[i], entity.Menu{Path: sub[i].Path}).Error; err != nil {
				return err
			}
		}

		// 3. 按钮级权限（每个子菜单 4 个）
		btns := []entity.Permission{}
		for _, m := range sub {
			prefix := m.Name // 如 User / UserGroup / Menu / Perm / OperLog
			btns = append(btns,
				entity.Permission{PermCode: *prefix + ":add", PermName: m.Title + "-新增", Sort: 1, Status: 1},
				entity.Permission{PermCode: *prefix + ":edit", PermName: m.Title + "-修改", Sort: 2, Status: 1},
				entity.Permission{PermCode: *prefix + ":del", PermName: m.Title + "-删除", Sort: 3, Status: 1},
				entity.Permission{PermCode: *prefix + ":view", PermName: m.Title + "-查看", Sort: 4, Status: 1},
			)
		}
		for i := range btns {
			if err := tx.FirstOrCreate(&btns[i], entity.Permission{PermCode: btns[i].PermCode}).Error; err != nil {
				return err
			}
		}

		// 4. 每个子菜单再挂一个「*」通配权限（方便前端整页授权）
		for _, m := range sub {
			all := entity.Permission{PermCode: *m.Name + ":*", PermName: m.Title + "-全部权限", Sort: 0, Status: 1}
			if err := tx.FirstOrCreate(&all, entity.Permission{PermCode: all.PermCode}).Error; err != nil {
				return err
			}
			// 5. 权限 <=> 菜单 绑定
			if err := tx.FirstOrCreate(&entity.PermissionMenuRel{PermID: all.ID, MenuID: m.ID},
				entity.PermissionMenuRel{PermID: all.ID, MenuID: m.ID}).Error; err != nil {
				return err
			}
		}

		logger.L.Info("menu & perm data seeded")
		return nil
	}); err != nil {
		logger.L.Error("seed menu perm fail", zap.Error(err))
	}
}

// 小工具
func stringPtr(s string) *string { return &s }
