CREATE TABLE IF NOT EXISTS oauth_user
(
    id         serial PRIMARY KEY,
    role_id    int,
    username   VARCHAR(50),
    password   VARCHAR(20),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);