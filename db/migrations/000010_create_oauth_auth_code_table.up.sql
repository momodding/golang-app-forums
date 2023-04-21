CREATE TABLE IF NOT EXISTS oauth_auth_code
(
    id           serial PRIMARY KEY,
    client_id    int,
    user_id      int,
    code         VARCHAR(50),
    redirect_uri VARCHAR(255),
    scope        VARCHAR(100),
    expires_at   TIMESTAMPTZ,
    created_at   TIMESTAMPTZ,
    updated_at   TIMESTAMPTZ,
    deleted_at   TIMESTAMPTZ
);