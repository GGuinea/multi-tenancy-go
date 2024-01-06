-- +goose Up
CREATE TABLE IF NOT EXISTS note (
	id text NOT NULL,
	title text NOT NULL,
	body text,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW(),
	CONSTRAINT tenant_pkey PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE note
