-- +goose Up
CREATE TABLE IF NOT EXISTS tenant (
	id text NOT NULL,
	name text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	active boolean NOT NULL DEFAULT true,
	CONSTRAINT tenant_pkey PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE tenant
