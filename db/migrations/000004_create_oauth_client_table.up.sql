CREATE TABLE IF NOT EXISTS oauth_client
(
    id           serial PRIMARY KEY,
    key          VARCHAR(255),
    secret       VARCHAR(100),
    redirect_uri VARCHAR(255),
    created_at   TIMESTAMPTZ,
    updated_at   TIMESTAMPTZ,
    deleted_at   TIMESTAMPTZ
);