CREATE TYPE "visibility" AS ENUM ('public', 'private', 'protected');

CREATE TABLE "snippets" (
    "id" BIGSERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "content" TEXT NOT NULL,
    "language" VARCHAR(10) NOT NULL,
    "user_id" BIGINT NOT NULL,
    "expires_at" timestamptz,
    "visibility" "visibility" NOT NULL DEFAULT 'public',
    "is_encrypted" BOOLEAN NOT NULL DEFAULT FALSE,
    "password_hash" TEXT,
    "view_count" BIGINT NOT NULL DEFAULT 0,
    "stars_count" BIGINT NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_snippets_id" ON "snippets" ("user_id");
CREATE INDEX "idx_snippets_language" ON "snippets" ("language");
CREATE INDEX "idx_snippets_visibility" ON "snippets" ("visibility");

ALTER TABLE "snippets" ADD CONSTRAINT "fk_users_snippets" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
