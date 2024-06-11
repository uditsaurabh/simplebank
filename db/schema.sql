CREATE TABLE IF NOT EXISTS accounts (
    id bigserial PRIMARY KEY,
    owner varchar NOT NULL,
    balance bigint NOT NULL,
    currency varchar NOT NULL,
    country_code bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE IF NOT EXISTS entries (
    id bigserial PRIMARY KEY,
    account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE transfers (
    id bigserial PRIMARY KEY,
    from_account_id bigint NOT NULL,
    to_account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "users" ( 
    username varchar PRIMARY KEY,
    hash_password varchar NOT NULL,
    full_name varchar NOT NULL,
    email varchar UNIQUE NOT NULL,
    password_changes_at timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    created_at timestamptz NOT NULL DEFAULT 'now()'
);