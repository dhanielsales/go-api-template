BEGIN;

CREATE TABLE IF NOT EXISTS product(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR (50) UNIQUE NOT NULL,
  description VARCHAR (300) NULL,
  price NUMERIC(10,2) NOT NULL,
  category_id UUID NOT NULL REFERENCES category(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NULL
);


COMMIT;
