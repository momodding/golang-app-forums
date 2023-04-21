CREATE TABLE IF NOT EXISTS post
(
    id         serial PRIMARY KEY,
    title      VARCHAR(50),
    post       VARCHAR(255),
    posted_by  int,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);