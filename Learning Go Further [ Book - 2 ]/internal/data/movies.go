package data

import (
	"database/sql"
	"encoding/json"
	"errors"
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
		plot = $7, language = $8, country = $9, awards = $10 WHERE id = $11
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
