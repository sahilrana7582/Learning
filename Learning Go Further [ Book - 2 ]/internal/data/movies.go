package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Movie struct {
	ID       int64    `json:"id"`
	Title    string   `json:"title"`
	Year     int      `json:"release_year"`
	Runtime  int      `json:"runtime"`
	Genre    []string `json:"genre"`
	Director string   `json:"director"`
	Actors   []string `json:"actors"`
	Plot     string   `json:"plot"`
	Language string   `json:"language"`
	Country  string   `json:"country"`
	Awards   string   `json:"awards,omitempty"`
}

type MovieModal struct {
	DB *sql.DB
}

func (m *MovieModal) Insert(movie *Movie) error {
	stmt := `
		INSERT INTO movies (title, release_year, runtime, genre, director, actors, plot, language, country, awards)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;
	`

	// Serialize Genre and Actors slices to JSON
	genreJSON, err := json.Marshal(movie.Genre)
	if err != nil {
		return err
	}
	actorsJSON, err := json.Marshal(movie.Actors)
	if err != nil {
		return err
	}

	err = m.DB.QueryRow(stmt, movie.Title, movie.Year, movie.Runtime, genreJSON, movie.Director, actorsJSON,
		movie.Plot, movie.Language, movie.Country, movie.Awards).Scan(&movie.ID)
	return err
}

func (m *MovieModal) Get(id int64) (*Movie, error) {
	stmt := `
		SELECT id, title, release_year, runtime, genre, director, actors, plot, language, country, awards
		FROM movies WHERE id = $1
	`

	row := m.DB.QueryRow(stmt, id)
	movie := &Movie{}

	// Intermediate variables for JSON data
	var genreJSON, actorsJSON []byte

	err := row.Scan(
		&movie.ID, &movie.Title, &movie.Year, &movie.Runtime,
		&genreJSON, &movie.Director, &actorsJSON,
		&movie.Plot, &movie.Language, &movie.Country, &movie.Awards,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // or custom ErrNotFound
		}
		return nil, err
	}

	// Deserialize JSON fields into slices
	if err = json.Unmarshal(genreJSON, &movie.Genre); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(actorsJSON, &movie.Actors); err != nil {
		return nil, err
	}

	return movie, nil
}

func (m *MovieModal) GetAll() ([]*Movie, error) {
	stmt := `
		SELECT id, title, release_year, runtime, genre, director, actors, plot, language, country, awards
		FROM movies
	`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		movie := &Movie{}
		var genreJSON, actorsJSON []byte

		err := rows.Scan(
			&movie.ID, &movie.Title, &movie.Year, &movie.Runtime,
			&genreJSON, &movie.Director, &actorsJSON,
			&movie.Plot, &movie.Language, &movie.Country, &movie.Awards,
		)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(genreJSON, &movie.Genre); err != nil {
			return nil, err
		}
		if err = json.Unmarshal(actorsJSON, &movie.Actors); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (m *MovieModal) Update(movie *Movie) error {
	stmt := `
		UPDATE movies SET title = $1, release_year = $2, runtime = $3, genre = $4, director = $5, actors = $6,
		plot = $7, language = $8, country = $9, awards = $10, version = version + 1 WHERE id = $11
		RETURNING version;
	`

	genreJSON, err := json.Marshal(movie.Genre)
	if err != nil {
		return err
	}
	actorsJSON, err := json.Marshal(movie.Actors)
	if err != nil {
		return err
	}

	_, err = m.DB.Exec(stmt, movie.Title, movie.Year, movie.Runtime, genreJSON, movie.Director, actorsJSON,
		movie.Plot, movie.Language, movie.Country, movie.Awards, movie.ID)
	return err
}

func (m *MovieModal) Delete(id int64) error {
	stmt := `DELETE FROM movies WHERE id = $1`
	_, err := m.DB.Exec(stmt, id)
	return err
}

func (m *MovieModal) GetAllMovieWithQuery(title string, genres []string, filters Filters) ([]*Movie, error) {
	query := `
		SELECT id, title, release_year, runtime, genre, director, actors, plot, language, country, awards
		FROM movies
		WHERE 1=1`

	args := []interface{}{}
	argID := 1

	// Case-insensitive title search
	if title != "" {
		query += fmt.Sprintf(" AND title ILIKE $%d", argID)
		args = append(args, "%"+title+"%")
		argID++
	}

	// Genre filter (assuming genre is a jsonb column)
	if len(genres) > 0 {
		query += fmt.Sprintf(`
			AND EXISTS (
				SELECT 1 FROM jsonb_array_elements_text(genre) AS g
				WHERE LOWER(g.value) = ANY($%d)
			)`, argID)

		for i, g := range genres {
			genres[i] = strings.ToLower(g)
		}
		args = append(args, pq.Array(genres))
		argID++
	}

	// Sorting (use safe validation on filters.Sort if it comes from user input)
	query += fmt.Sprintf(" ORDER BY %s", filters.Sort)

	// Pagination
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, filters.Limit(), filters.Offset())

	// Database query with context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		var genreJSON, actorsJSON []byte

		if err := rows.Scan(
			&movie.ID, &movie.Title, &movie.Year, &movie.Runtime,
			&genreJSON, &movie.Director, &actorsJSON,
			&movie.Plot, &movie.Language, &movie.Country, &movie.Awards,
		); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}

		if err := json.Unmarshal(genreJSON, &movie.Genre); err != nil {
			return nil, fmt.Errorf("genre unmarshal error: %w", err)
		}

		if err := json.Unmarshal(actorsJSON, &movie.Actors); err != nil {
			return nil, fmt.Errorf("actors unmarshal error: %w", err)
		}

		movies = append(movies, &movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return movies, nil
}
