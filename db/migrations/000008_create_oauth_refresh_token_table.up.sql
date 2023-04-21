CREATE TABLE IF NOT EXISTS oauth_refresh_token
(
    id         serial PRIMARY KEY,
    client_id  int,
    user_id    int,
    token      VARCHAR(255),
    scope      VARCHAR(100),
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);