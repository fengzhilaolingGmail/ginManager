-- 1. 用户
CREATE TABLE `user` (
    id            BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    username      VARCHAR(50)  NOT NULL UNIQUE COMMENT '登录账号',
    password_hash CHAR(60)     NOT NULL COMMENT 'bcrypt 哈希',
    nickname      VARCHAR(50)  NOT NULL COMMENT '显示名',
    email         VARCHAR(100) UNIQUE COMMENT '邮箱',
    phone         VARCHAR(20)  UNIQUE COMMENT '手机',
    status        TINYINT      NOT NULL DEFAULT 1 COMMENT '1=启用 0=禁用',
    created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户';

-- 2. 用户组
CREATE TABLE `user_group` (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    group_code  VARCHAR(50) NOT NULL UNIQUE COMMENT '组编码',
    group_name  VARCHAR(50) NOT NULL COMMENT '组名称',
    sort        INT         NOT NULL DEFAULT 0 COMMENT '排序',
    status      TINYINT     NOT NULL DEFAULT 1 COMMENT '1=启用 0=禁用',
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户组';

-- 3. 用户 ↔ 用户组
CREATE TABLE `user_group_rel` (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    user_id    BIGINT NOT NULL COMMENT '用户ID',
    group_id   BIGINT NOT NULL COMMENT '组ID',
    UNIQUE KEY uk_user_group (user_id, group_id),
    CONSTRAINT fk_ug_user  FOREIGN KEY (user_id)  REFERENCES `user`(id)        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_ug_group FOREIGN KEY (group_id) REFERENCES `user_group`(id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户与组关联';

-- 4. 角色
CREATE TABLE `role` (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    role_code   VARCHAR(50) NOT NULL UNIQUE COMMENT '角色编码',
    role_name   VARCHAR(50) NOT NULL COMMENT '角色名称',
    sort        INT         NOT NULL DEFAULT 0 COMMENT '排序',
    status      TINYINT     NOT NULL DEFAULT 1 COMMENT '1=启用 0=禁用',
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色';

-- 5. 用户组 ↔ 角色
CREATE TABLE `group_role_rel` (
    id        BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    group_id  BIGINT NOT NULL COMMENT '组ID',
    role_id   BIGINT NOT NULL COMMENT '角色ID',
    UNIQUE KEY uk_group_role (group_id, role_id),
    CONSTRAINT fk_gr_group FOREIGN KEY (group_id) REFERENCES `user_group`(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_gr_role  FOREIGN KEY (role_id)  REFERENCES `role`(id)       ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='组与角色关联';

-- 6. 权限
CREATE TABLE `permission` (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    perm_code   VARCHAR(100) NOT NULL UNIQUE COMMENT '权限编码，如 system:user:add',
    perm_name   VARCHAR(100) NOT NULL COMMENT '权限名称',
    sort        INT          NOT NULL DEFAULT 0 COMMENT '排序',
    status      TINYINT      NOT NULL DEFAULT 1 COMMENT '1=启用 0=禁用',
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限';

-- 7. 角色 ↔ 权限
CREATE TABLE `role_permission_rel` (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    role_id    BIGINT NOT NULL COMMENT '角色ID',
    perm_id    BIGINT NOT NULL COMMENT '权限ID',
    UNIQUE KEY uk_role_perm (role_id, perm_id),
    CONSTRAINT fk_rp_role FOREIGN KEY (role_id) REFERENCES `role`(id)           ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_rp_perm FOREIGN KEY (perm_id) REFERENCES `permission`(id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色与权限关联';

-- 8. 菜单
CREATE TABLE `menu` (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    parent_id  BIGINT DEFAULT 0 COMMENT '父菜单ID，0=根',
    path       VARCHAR(200) COMMENT '前端路由',
    component  VARCHAR(100) COMMENT '前端组件路径',
    name       VARCHAR(50) COMMENT '路由名称',
    title      VARCHAR(50) NOT NULL COMMENT '菜单标题',
    icon       VARCHAR(50) COMMENT '图标',
    sort       INT         NOT NULL DEFAULT 0 COMMENT '排序',
    status     TINYINT     NOT NULL DEFAULT 1 COMMENT '1=启用 0=禁用',
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单';

-- 9. 权限 ↔ 菜单
CREATE TABLE `permission_menu_rel` (
    id        BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    perm_id   BIGINT NOT NULL COMMENT '权限ID',
    menu_id   BIGINT NOT NULL COMMENT '菜单ID',
    UNIQUE KEY uk_perm_menu (perm_id, menu_id),
    CONSTRAINT fk_pm_perm FOREIGN KEY (perm_id) REFERENCES `permission`(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_pm_menu FOREIGN KEY (menu_id) REFERENCES `menu`(id)       ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限与菜单关联';

-- 10. 用户操作日志
CREATE TABLE `user_log` (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    user_id     BIGINT COMMENT '用户ID',
    module      VARCHAR(50)  COMMENT '模块',
    action      VARCHAR(50)  COMMENT '动作',
    method      VARCHAR(10)  COMMENT 'HTTP方法',
    path        VARCHAR(200) COMMENT '请求路径',
    ip          VARCHAR(45)  COMMENT '客户端IP',
    user_agent  TEXT         COMMENT 'UA',
    status      TINYINT      NOT NULL DEFAULT 1 COMMENT '1=成功 0=失败',
    error_msg   TEXT         COMMENT '错误信息',
    duration_ms INT          COMMENT '耗时(ms)',
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户操作日志';