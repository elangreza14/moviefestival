BEGIN
;

CREATE TABLE IF NOT EXISTS "artists" (
    "name" VARCHAR PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_artist_update" BEFORE
UPDATE
    ON "artists" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

CREATE TABLE IF NOT EXISTS "movie_artists" (
    "artist_name" VARCHAR REFERENCES artists(name),
    "movie_id" BIGINT REFERENCES movies(id),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("artist_name", "movie_id")
);

COMMIT;