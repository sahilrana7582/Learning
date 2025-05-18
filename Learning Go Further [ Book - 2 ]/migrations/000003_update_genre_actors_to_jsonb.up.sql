-- Up migration: convert genre and actors from text[] to jsonb

ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_genre_check;
ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_actors_check;

ALTER TABLE movies
  ALTER COLUMN genre TYPE jsonb USING genre::text::jsonb,
  ALTER COLUMN actors TYPE jsonb USING actors::text::jsonb;

ALTER TABLE movies ADD CONSTRAINT movies_genre_check CHECK (jsonb_array_length(genre) > 0);
ALTER TABLE movies ADD CONSTRAINT movies_actors_check CHECK (jsonb_array_length(actors) > 0);
