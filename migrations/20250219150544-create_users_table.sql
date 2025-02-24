
-- +migrate Up
    CREATE TABLE IF NOT EXISTS users (
        id bigserial NOT NULL,
        username VARCHAR(200) NOT NULL,
        password VARCHAR(200) NOT NULL,
        email VARCHAR NOT NULL,
        role_id INTEGER NOT NULL,
        branch_id INTEGER,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT user_pkey PRIMARY KEY (id)
    )

    CREATE TABLE IF NOT EXISTS roles (
        id bigserial NOT NULL,
        code VARCHAR NOT NULL,
        name VARCHAR,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT roles_pkey PRIMARY KEY (id)
    )

    CREATE TABLE IF NOT EXISTS permissions (
        id bigserial NOT NULL,
        code VARCHAR NOT NULL,
        name VARCHAR,
        action VARCHAR,LL
        group_menu VARCHAR,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT permissions_pkey PRIMARY KEY (id)
    )

    CREATE TABLE IF NOT EXISTS role_permissions (
        id bigserial NOT NULL,
        role_id INTEGER,
        permissions_id INTEGER,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT roles_permissions_pkey PRIMARY KEY (id)
    )

-- +migrate Down
