CREATE TABLE IF NOT EXISTS post_comment
(
    id                serial PRIMARY KEY,
    post_id           int,
    parent_comment_id int,
    content           VARCHAR(255),
    posted_by         int,
    created_at        TIMESTAMPTZ,
    updated_at        TIMESTAMPTZ,
    deleted_at        TIMESTAMPTZ
);