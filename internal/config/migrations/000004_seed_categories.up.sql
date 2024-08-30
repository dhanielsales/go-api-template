BEGIN;

INSERT INTO category ("id", "name", "slug", "description", "created_at")
VALUES ('8bcb62db-66a5-446b-9a95-9ae0f72e31a5', 'Electronics', 'electronics', 'Electronics', 1706208670948);

INSERT INTO category ("id", "name", "slug", "description", "created_at")
VALUES ('eb0b97c4-f466-4415-a2ee-e664c56e97e8', 'Books', 'books', 'Books', 1706208670948);

INSERT INTO category ("id", "name", "slug", "description", "created_at")
VALUES ('64655489-ea4d-433a-aed5-a4dce06b87a0', 'Clothing', 'clothing', 'Clothing', 1706208670948);

COMMIT;
