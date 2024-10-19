BEGIN;

DROP INDEX IF EXISTS "users_id_index";

DROP INDEX IF EXISTS "users_username_index";

DROP INDEX IF EXISTS "users_email_index";

DROP TABLE IF EXISTS "users" CASCADE;

COMMIT;