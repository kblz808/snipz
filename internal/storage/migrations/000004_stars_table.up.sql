CREATE TABLE "stars" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "snippet_id" BIGINT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_stars_snippet_id" ON "stars" ("snippet_id");
CREATE INDEX "idx_stars_user_id" ON "stars" ("user_id");

ALTER TABLE "stars" ADD CONSTRAINT "fk_stars_snippets" FOREIGN KEY ("snippet_id") REFERENCES "snippets" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "stars" ADD CONSTRAINT "fk_stars_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "stars" ADD CONSTRAINT "unique_stars_snippet_user" UNIQUE ("snippet_id", "user_id");
