BEGIN
;

-- title varchar
-- description text
-- duration bigint
-- watch_url varchar
-- artists, make new table
-- genres, make new table
CREATE TABLE IF NOT EXISTS "genres" (
    "name" VARCHAR PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_genre_update" BEFORE
UPDATE
    ON "genres" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

CREATE TABLE IF NOT EXISTS "movie_genres" (
    "genre_name" VARCHAR REFERENCES genres(name),
    "movie_id" BIGINT REFERENCES movies(id),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("genre_name", "movie_id")
);

COMMIT;