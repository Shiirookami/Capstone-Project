BEGIN;
CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255);
    "password" VARCHAR(255) NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6)
);
COMMIT;