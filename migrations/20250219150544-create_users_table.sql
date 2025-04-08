
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
    )

    CREATE TABLE IF NOT EXISTS roles (
        id bigserial NOT NULL,
        code VARCHAR NOT NULL,
        name VARCHAR,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by INTEGER,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_by INTEGER,
        CONSTRAINT roles_pkey PRIMARY KEY (id)
    )

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
    )

    CREATE TABLE IF NOT EXISTS role_permissions (
        id bigserial NOT NULL,
        role_id INTEGER,
        permissions_id INTEGER,
        access_scope VARCHAR CHECK (access_scope IN ('own', 'all')) DEFAULT 'own',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by INTEGER,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_by INTEGER,
        CONSTRAINT roles_permissions_pkey PRIMARY KEY (id)
    )

-- +migrate Down
