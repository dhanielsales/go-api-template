BEGIN;

CREATE TABLE IF NOT EXISTS product(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR (50) UNIQUE NOT NULL,
  slug VARCHAR (50) UNIQUE NOT NULL,
  description VARCHAR (300) NULL,
  price DECIMAL CHECK (price > 0) NOT NULL,
  category_id UUID NOT NULL REFERENCES category(id),
  created_at bigint NOT NULL,
  updated_at bigint NULL
);


COMMIT;
