-- +goose Up
CREATE TABLE shoots (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    location TEXT NOT NULL,
    camera TEXT NOT NULL,
    lens TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('planned', 'shot', 'edited', 'published', 'archived')),
    shoot_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE shoots;
