package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"version"`
}

type password struct {
	plaintext *string
	hash      []byte
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

func (p *password) Set(passwordText string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordText), 4)

	if err != nil {
		return err
	}

	p.plaintext = &passwordText
	p.hash = hash
	return nil
}

func (p *password) Compare(passwordText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(passwordText))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(user *User) error {
	stmt := `
		INSERT INTO users (name, email, password_hash, activated)
		VALUES ($1, $2, $3, $4)
		RETURNING id, version;
	`

	args := []interface{}{
		user.Name,
		user.Email,
		user.Password,
		user.Activated,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, stmt, args...).Scan(&user.ID, &user.Version)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			if pqErr.Constraint == "users_email_key" {
				return fmt.Errorf("email already exists")
			}
		}
		return err
	}

	return nil
}

func (u *UserModel) GetByEmail(email string) (*User, error) {
	stmt := `
		SELECT id, created_at, name, email, password_hash, activated, version
		FROM users
		WHERE email = $1;
	`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, stmt, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Activated,
		&user.Version,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no record found for email: %s", email)
		}
		return nil, err
	}

	return &user, nil
}
