-- rollback: 20250518_remove_version_column_from_movies.sql

ALTER TABLE movies
DROP COLUMN version;
