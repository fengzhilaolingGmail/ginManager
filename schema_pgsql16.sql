-- 统一触发器：自动更新 updated_at
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 1. 用户
CREATE TABLE "user" (
    id            BIGSERIAL PRIMARY KEY,
    username      VARCHAR(50)  NOT NULL UNIQUE,
    password_hash CHAR(60)     NOT NULL,
    nickname      VARCHAR(50)  NOT NULL,
    email         VARCHAR(100) UNIQUE,
    phone         VARCHAR(20)  UNIQUE,
    status        SMALLINT     NOT NULL DEFAULT 1,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE "user" IS '用户';
COMMENT ON COLUMN "user".status IS '1=启用 0=禁用';
CREATE TRIGGER trg_user_updated BEFORE UPDATE ON "user"
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- 2. 用户组
CREATE TABLE user_group (
    id          BIGSERIAL PRIMARY KEY,
    group_code  VARCHAR(50) NOT NULL UNIQUE,
    group_name  VARCHAR(50) NOT NULL,
    sort        INT         NOT NULL DEFAULT 0,
    status      SMALLINT    NOT NULL DEFAULT 1,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE user_group IS '用户组';
CREATE TRIGGER trg_group_updated BEFORE UPDATE ON user_group
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- 3. 用户 ↔ 用户组
CREATE TABLE user_group_rel (
    id       BIGSERIAL PRIMARY KEY,
    user_id  BIGINT NOT NULL REFERENCES "user"(id)        ON DELETE CASCADE ON UPDATE CASCADE,
    group_id BIGINT NOT NULL REFERENCES user_group(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (user_id, group_id)
);
COMMENT ON TABLE user_group_rel IS '用户与组关联';

-- 4. 角色
CREATE TABLE role (
    id          BIGSERIAL PRIMARY KEY,
    role_code   VARCHAR(50) NOT NULL UNIQUE,
    role_name   VARCHAR(50) NOT NULL,
    sort        INT         NOT NULL DEFAULT 0,
    status      SMALLINT    NOT NULL DEFAULT 1,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE role IS '角色';
CREATE TRIGGER trg_role_updated BEFORE UPDATE ON role
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- 5. 用户组 ↔ 角色
CREATE TABLE group_role_rel (
    id       BIGSERIAL PRIMARY KEY,
    group_id BIGINT NOT NULL REFERENCES user_group(id) ON DELETE CASCADE ON UPDATE CASCADE,
    role_id  BIGINT NOT NULL REFERENCES role(id)       ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (group_id, role_id)
);
COMMENT ON TABLE group_role_rel IS '组与角色关联';

-- 6. 权限
CREATE TABLE permission (
    id          BIGSERIAL PRIMARY KEY,
    perm_code   VARCHAR(100) NOT NULL UNIQUE,
    perm_name   VARCHAR(100) NOT NULL,
    sort        INT          NOT NULL DEFAULT 0,
    status      SMALLINT     NOT NULL DEFAULT 1,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE permission IS '权限';
CREATE TRIGGER trg_perm_updated BEFORE UPDATE ON permission
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- 7. 角色 ↔ 权限
CREATE TABLE role_permission_rel (
    id      BIGSERIAL PRIMARY KEY,
    role_id BIGINT NOT NULL REFERENCES role(id)       ON DELETE CASCADE ON UPDATE CASCADE,
    perm_id BIGINT NOT NULL REFERENCES permission(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (role_id, perm_id)
);
COMMENT ON TABLE role_permission_rel IS '角色与权限关联';

-- 8. 菜单
CREATE TABLE menu (
    id         BIGSERIAL PRIMARY KEY,
    parent_id  BIGINT DEFAULT 0,
    path       VARCHAR(200),
    component  VARCHAR(100),
    name       VARCHAR(50),
    title      VARCHAR(50) NOT NULL,
    icon       VARCHAR(50),
    sort       INT         NOT NULL DEFAULT 0,
    status     SMALLINT    NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE menu IS '菜单';
CREATE TRIGGER trg_menu_updated BEFORE UPDATE ON menu
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- 9. 权限 ↔ 菜单
CREATE TABLE permission_menu_rel (
    id      BIGSERIAL PRIMARY KEY,
    perm_id BIGINT NOT NULL REFERENCES permission(id) ON DELETE CASCADE ON UPDATE CASCADE,
    menu_id BIGINT NOT NULL REFERENCES menu(id)       ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (perm_id, menu_id)
);
COMMENT ON TABLE permission_menu_rel IS '权限与菜单关联';

-- 10. 用户操作日志
CREATE TABLE user_log (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT,
    module      VARCHAR(50),
    action      VARCHAR(50),
    method      VARCHAR(10),
    path        VARCHAR(200),
    ip          VARCHAR(45),
    user_agent  TEXT,
    status      SMALLINT     NOT NULL DEFAULT 1,
    error_msg   TEXT,
    duration_ms INT,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE user_log IS '用户操作日志';