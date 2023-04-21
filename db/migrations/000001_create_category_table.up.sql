CREATE TABLE IF NOT EXISTS category
(
    id
    serial
    PRIMARY
    KEY,
    name
    VARCHAR
(
    50
),
    description VARCHAR
(
    255
),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
    );