begin;

drop trigger if exists set_updated_at_timestamp_users_table on "users";
drop trigger if exists set_updated_at_timestamp_oauth_table on "oauth";
drop trigger if exists  set_updated_at_timestamp_products_table on "products";
drop trigger if exists  set_updated_at_timestamp_images_table on "images";
drop trigger if exists  set_updated_at_timestamp_orders_table on "orders";

drop function if exists set_updated_at_column();

drop table if exists "users" cascade;
drop table if exists "oauth" cascade;
drop table if exists "roles" cascade;
drop table if exists "products" cascade;
drop table if exists "products_categories" cascade;
drop table if exists "categories" cascade;
drop table if exists "images" cascade;
drop table if exists "orders" cascade;
drop table if exists "products_orders" cascade;


drop sequence if exists users_id_sequence;
drop sequence if exists products_id_sequence;
drop sequence if exists orders_id_sequence;


drop type if exists "order_status";

commit;