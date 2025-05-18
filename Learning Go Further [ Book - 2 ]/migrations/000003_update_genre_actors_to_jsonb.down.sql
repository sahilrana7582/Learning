-- Down migration: revert genre and actors from jsonb to text[]

ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_genre_check;
ALTER TABLE movies DROP CONSTRAINT IF EXISTS movies_actors_check;

ALTER TABLE movies
  ALTER COLUMN genre TYPE text[] USING genre::text::text[],
  ALTER COLUMN actors TYPE text[] USING actors::text::text[];

ALTER TABLE movies ADD CONSTRAINT movies_genre_check CHECK (array_length(genre, 1) > 0);
ALTER TABLE movies ADD CONSTRAINT movies_actors_check CHECK (array_length(actors, 1) > 0);
