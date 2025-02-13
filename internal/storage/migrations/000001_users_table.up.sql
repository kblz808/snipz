CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL,
    "hashed_password" TEXT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_users_username" ON "users" ("username");

ALTER TABLE "users" ADD CONSTRAINT "unique_users_username" UNIQUE ("username");
