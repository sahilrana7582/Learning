ALTER TABLE movies ADD CONSTRAINT movies_release_year_check CHECK (release_year > 1888);
ALTER TABLE movies ADD CONSTRAINT movies_runtime_check CHECK (runtime > 0);
ALTER TABLE movies ADD CONSTRAINT movies_genre_check CHECK (array_length(genre, 1) > 0);
ALTER TABLE movies ADD CONSTRAINT movies_actors_check CHECK (array_length(actors, 1) > 0);
ALTER TABLE movies ADD CONSTRAINT movies_language_check CHECK (language != '');
ALTER TABLE movies ADD CONSTRAINT movies_country_check CHECK (country != '');
ALTER TABLE movies ADD CONSTRAINT movies_awards_check CHECK (awards != '');
ALTER TABLE movies ADD CONSTRAINT movies_title_check CHECK (title != '');
ALTER TABLE movies ADD CONSTRAINT movies_director_check CHECK (director != '');