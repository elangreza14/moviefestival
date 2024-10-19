BEGIN
;

-- title varchar
-- description text
-- duration bigint
-- watch_url varchar
-- artists, make new table
-- genres, make new table
CREATE TABLE IF NOT EXISTS "movies" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR,
    "description" TEXT,
    "watch_url" VARCHAR,
    "duration" INTERVAL NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_movie_update" BEFORE
UPDATE
    ON "movies" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;