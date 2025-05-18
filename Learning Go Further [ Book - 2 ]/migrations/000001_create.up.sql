CREATE TABLE IF NOT EXISTS movies (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    release_year integer NOT NULL,
    runtime integer NOT NULL,
    genre text[] NOT NULL,
    director text NOT NULL,
    actors text[] NOT NULL,
    plot text NOT NULL,
    language text NOT NULL,
    country text NOT NULL,
    awards text
);
