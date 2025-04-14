-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id bigserial NOT NULL,
    name VARCHAR(200) NOT NULL,
    username VARCHAR(200) NOT NULL,
    password VARCHAR(200) NOT NULL,
    email VARCHAR NOT NULL,
    role_id INTEGER NOT NULL,
    branch_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER,
    CONSTRAINT user_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS roles (
    id bigserial NOT NULL,
    code VARCHAR NOT NULL,
    name VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER,
    CONSTRAINT roles_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS permissions (
    id bigserial NOT NULL,
    code VARCHAR NOT NULL,
    name VARCHAR,
    group_menu VARCHAR,
    action VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER,
    CONSTRAINT permissions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS role_permissions (
    id bigserial NOT NULL,
    role_id INTEGER NOT NULL,
    permissions_id INTEGER NOT NULL,
    access_scope VARCHAR CHECK (access_scope IN ('own', 'all')) DEFAULT 'own',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER,
    CONSTRAINT role_permissions_pkey PRIMARY KEY (id),
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles (id),
    CONSTRAINT fk_permission FOREIGN KEY (permissions_id) REFERENCES permissions (id)
);

INSERT INTO permissions (code, name, group_menu, action, created_by, updated_by) VALUES
-- User
('user:create', 'Permission to create user data (user-create)', 'user', 'create', 1, 1),
('user:read', 'Permission to read user data (user-read)', 'user', 'read', 1, 1),
('user:update', 'Permission to update user data (user-update)', 'user', 'update', 1, 1),
('user:delete', 'Permission to delete user data (user-delete)', 'user', 'delete', 1, 1),

-- Category
('category:create', 'Permission to create category data (category-create)', 'category', 'create', 1, 1),
('category:read', 'Permission to read category data (category-read)', 'category', 'read', 1, 1),
('category:update', 'Permission to update category data (category-update)', 'category', 'update', 1, 1),
('category:delete', 'Permission to delete category data (category-delete)', 'category', 'delete', 1, 1),

-- Role
('role:create', 'Permission to create role data (role-create)', 'role', 'create', 1, 1),
('role:read', 'Permission to read role data (role-read)', 'role', 'read', 1, 1),
('role:update', 'Permission to update role data (role-update)', 'role', 'update', 1, 1),
('role:delete', 'Permission to delete role data (role-delete)', 'role', 'delete', 1, 1),

-- Permissions
('permissions:create', 'Permission to create permissions data (permissions-create)', 'permissions', 'create', 1, 1),
('permissions:read', 'Permission to read permissions data (permissions-read)', 'permissions', 'read', 1, 1),
('permissions:update', 'Permission to update permissions data (permissions-update)', 'permissions', 'update', 1, 1),
('permissions:delete', 'Permission to delete permissions data (permissions-delete)', 'permissions', 'delete', 1, 1),

-- Role Permissions
('role_permissions:create', 'Permission to create role permissions data (role_permissions-create)', 'role_permissions', 'create', 1, 1),
('role_permissions:read', 'Permission to read role permissions data (role_permissions-read)', 'role_permissions', 'read', 1, 1),
('role_permissions:update', 'Permission to update role permissions data (role_permissions-update)', 'role_permissions', 'update', 1, 1),
('role_permissions:delete', 'Permission to delete role permissions data (role_permissions-delete)', 'role_permissions', 'delete', 1, 1);

-- +migrate Down
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;