CREATE TABLE IF NOT EXISTS oauth_role
(
    id         serial PRIMARY KEY,
    name       VARCHAR(50),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);