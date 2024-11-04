BEGIN;

TRUNCATE TABLE IF EXISTS "users" CASCADE;
TRUNCATE TABLE IF EXISTS "oauth" CASCADE;
TRUNCATE TABLE IF EXISTS "roles" CASCADE;
TRUNCATE TABLE IF EXISTS "products" CASCADE;
TRUNCATE TABLE IF EXISTS "products_categories" CASCADE;
TRUNCATE TABLE IF EXISTS "categories" CASCADE;
TRUNCATE TABLE IF EXISTS "images" CASCADE;
TRUNCATE TABLE IF EXISTS "orders" CASCADE;
TRUNCATE TABLE IF EXISTS "products_orders" CASCADE;

SELECT setval(pg_get_serial_sequence('"roles"', 'id'), 1, false);
SELECT setval(pg_get_serial_sequence('"categories"', 'id'), 1, false);


COMMIT;
