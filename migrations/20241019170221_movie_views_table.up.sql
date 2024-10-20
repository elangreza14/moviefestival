BEGIN
;

CREATE TABLE IF NOT EXISTS "movie_views" (
    "movie_id" BIGINT REFERENCES movies(id),
    "views" BIGINT DEFAULT 0,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL,
    PRIMARY KEY ("movie_id")
);

CREATE TRIGGER "log_movie_view_update" BEFORE
UPDATE
    ON "movie_views" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;