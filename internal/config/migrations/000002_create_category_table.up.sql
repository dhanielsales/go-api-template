BEGIN;

CREATE TABLE IF NOT EXISTS category(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR (50) NOT NULL,
  slug VARCHAR (50) UNIQUE NOT NULL,
  description VARCHAR (300) NULL,
  created_at bigint NOT NULL,
  updated_at bigint NULL
);

COMMIT;
