CREATE EXTENSION IF NOT EXISTS citext;
DROP TABLE IF EXISTS chapters;
DROP TABLE IF EXISTS gnovels;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS users_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS gnovels (
    id BIGSERIAL PRIMARY KEY,
    gntype TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    genres TEXT[] NOT NULL,
    status TEXT NOT NULL,
    nchapters INT NOT NULL,
    author TEXT NOT NULL,
    year INT NOT NULL,
    created_at timestamp(0) with time zone not null  default NOW(),
    UNIQUE(title)
);


CREATE TABLE IF NOT EXISTS chapters (
    id BIGSERIAL PRIMARY KEY,
    gnovel_id BIGINT NOT NULL REFERENCES gnovels(id) ON DELETE CASCADE,
    number INT NOT NULL,
    file_path TEXT NOT NULL,
    created_at timestamp(0) with time zone not null  default NOW()
);


CREATE TABLE IF NOT EXISTS users (
	id bigserial PRIMARY KEY,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	name text NOT NULL,
	email citext UNIQUE NOT NULL,
	password_hash bytea NOT NULL,
	activated bool NOT NULL
);


CREATE TABLE IF NOT EXISTS tokens (
	hash bytea PRIMARY KEY,
	user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
	expiry timestamp(0) with time zone NOT NULL,
	scope text NOT NULL
);



CREATE TABLE IF NOT EXISTS permissions (
	id bigserial PRIMARY KEY,
	code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
	user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
	permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
	PRIMARY KEY (user_id, permission_id)
);

INSERT INTO permissions (code)
VALUES
('gnovels:read'),
('gnovels:write');