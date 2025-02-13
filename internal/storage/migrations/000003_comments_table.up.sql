CREATE TABLE "comments" (
    "id" BIGSERIAL PRIMARY KEY,
    "snippet_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_comments_snippet_id" ON "comments" ("snippet_id");
CREATE INDEX "idx_comments_user_id" ON "comments" ("user_id");

ALTER TABLE "comments" ADD CONSTRAINT "fk_comments_snippets" FOREIGN KEY ("snippet_id") REFERENCES "snippets" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "comments" ADD CONSTRAINT "fk_comments_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
