CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "hash_password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changes_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "accounts"
ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

--CREATE UNIQUE INDEX ON "account" ("owner", "currency");docker compose 
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner","currency")