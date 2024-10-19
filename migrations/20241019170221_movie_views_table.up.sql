BEGIN
;

CREATE TABLE IF NOT EXISTS "movie_views" (
    "movie_id" BIGINT REFERENCES movies(id),
    "views" BIGINT DEFAULT 0,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("movie_id")
);

COMMIT;