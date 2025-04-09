-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id bigserial NOT NULL,
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
('user:write', 'Permission to write user data (user-write)', 'user', 'write', 1, 1),
('user:read', 'Permission to read user data (user-read)', 'user', 'read', 1, 1),
('category:write', 'Permission to write category data (category-write)', 'category', 'write', 1, 1),
('category:read', 'Permission to read category data (category-read)', 'category', 'read', 1, 1),
('role:write', 'Permission to write role data (role-write)', 'role', 'write', 1, 1),
('role:read', 'Permission to read role data (role-read)', 'role', 'read', 1, 1),
('permissions:write', 'Permission to write permissions data (permissions-write)', 'permissions', 'write', 1, 1),
('permissions:read', 'Permission to read permissions data (permissions-read)', 'permissions', 'read', 1, 1),
('role_permissions:read', 'Permission to read role permissions data (role_permissions-read)', 'role_permissions', 'read', 1, 1),
('role_permissions:write', 'Permission to write role permissions data (role_permissions-write)', 'role_permissions', 'write', 1, 1);

-- +migrate Down
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;