BEGIN;

SET TIME zone 'Asia/Bangkok';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE users_id_sequence START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE products_id_sequence START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE orders_id_sequence START WITH 1 INCREMENT BY 1;

CREATE TYPE "order_status" AS ENUM (
    'waiting',
    'shipping',
    'completed',
    'canceled'
);

CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TABLE "roles" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR NOT NULL UNIQUE
);


CREATE TABLE "users" (
    "id" VARCHAR PRIMARY KEY DEFAULT CONCAT('U', LPAD(NEXTVAL('users_id_sequence')::text, 6, '0')),
    "username" VARCHAR UNIQUE NOT NULL,
    "password" VARCHAR NOT NULL,
    "email" VARCHAR UNIQUE NOT NULL,
    "role_id" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "oauth" (
    "id" UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    "user_id" VARCHAR NOT NULL,
    "access_token" VARCHAR NOT NULL,
    "refresh_token" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "products" (
    "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('P', LPAD(NEXTVAL('products_id_sequence')::text, 6, '0')),
    "title" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL DEFAULT '',
    "price" FLOAT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "categories" (
    "id" SERIAL PRIMARY KEY NOT NULL UNIQUE,
    "title" VARCHAR UNIQUE NOT NULL
);

CREATE TABLE "products_categories" (
    "id" UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    "product_id" VARCHAR NOT NULL,
    "category_id" INT NOT NULL
);

CREATE TABLE "images" (
    "id" UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    "filename" VARCHAR NOT NULL,
    "url" VARCHAR NOT NULL,
    "product_id" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "orders" (
    "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('O', LPAD(NEXTVAL('orders_id_sequence')::text, 6, '0')),
    "user_id" VARCHAR NOT NULL,
    "contact" VARCHAR NOT NULL,
    "address" VARCHAR NOT NULL,
    "transfer_slip" JSONB,
    "status" order_status NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "products_orders" (
    "id" UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    "order_id" VARCHAR NOT NULL,
    "qty" INT DEFAULT 1,
    "product" JSONB
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "products_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

CREATE TRIGGER set_updated_at_timestamp_users_table BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_oauth_table BEFORE UPDATE ON "oauth" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_products_table BEFORE UPDATE ON "products" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_images_table BEFORE UPDATE ON "images" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_orders_table BEFORE UPDATE ON "orders" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();

COMMIT;
