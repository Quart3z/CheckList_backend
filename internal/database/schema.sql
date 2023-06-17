CREATE TABLE IF NOT EXISTS TaskCategory (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title text NOT NULL,
    is_default boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

ALTER TABLE TaskCategory ENABLE ROW LEVEL SECURITY;

CREATE TABLE IF NOT EXISTS Task (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title text NOT NULL,
    description text NOT NULL,
    is_completed boolean NOT NULL DEFAULT false,
    category_id uuid NOT NULL REFERENCES TaskCategory(id),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

ALTER TABLE Task ENABLE ROW LEVEL SECURITY;