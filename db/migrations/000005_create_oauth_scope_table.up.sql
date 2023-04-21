CREATE TABLE IF NOT EXISTS oauth_scope
(
    id          serial PRIMARY KEY,
    scope       VARCHAR(100),
    description VARCHAR(200),
    is_default  BOOLEAN,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    deleted_at  TIMESTAMPTZ
);