BEGIN
;

CREATE TABLE IF NOT EXISTS "movie_views_histories" (
    "id" SERIAL PRIMARY KEY,
    "movie_id" BIGINT REFERENCES movies(id),
    "user_id" UUID REFERENCES users(id),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

COMMIT;