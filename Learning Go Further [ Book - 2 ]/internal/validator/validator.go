package validator

import (
	"regexp"
	"strings"

	"github.com/sahilrana7582/Learning/internal/data"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, ok := v.Errors[key]; !ok {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MinLength(value string, min int) bool {
	return len(strings.TrimSpace(value)) >= min
}

func MaxLength(value string, max int) bool {
	return len(strings.TrimSpace(value)) <= max
}

func Matches(value string, pattern *regexp.Regexp) bool {
	return pattern.MatchString(value)
}

func NotZero[T comparable](value T) bool {
	var zero T
	return value != zero
}

func IsValidGenre(genre string) bool {
	validGenres := map[string]struct{}{
		"Action": {}, "Comedy": {}, "Drama": {}, "Horror": {}, "Sci-Fi": {},
	}
	_, ok := validGenres[genre]
	return ok
}

func NotDuplicateGenre(genres []string) bool {
	seen := make(map[string]bool)
	for _, genre := range genres {
		if seen[genre] {
			return false
		}
		seen[genre] = true
	}
	return true
}

func IsValidYear(year int) bool {
	return year > 1888 && year <= 2025
}

func IsValidRuntime(runtime int) bool {
	return runtime > 0 && runtime <= 300
}

func IsValidLanguage(language string) bool {
	validLanguages := map[string]struct{}{
		"English": {}, "Spanish": {}, "French": {}, "German": {}, "Chinese": {},
	}
	_, ok := validLanguages[language]
	return ok
}

func IsValidCountry(country string) bool {
	validCountries := map[string]struct{}{
		"USA": {}, "UK": {}, "Canada": {}, "Australia": {}, "India": {},
	}
	_, ok := validCountries[country]
	return ok
}

func AllValidGenres(genres []string) bool {
	for _, g := range genres {
		if !IsValidGenre(g) {
			return false
		}
	}
	return true
}

func AllNotBlank(items []string) bool {
	for _, item := range items {
		if !NotBlank(item) {
			return false
		}
	}
	return true
}

func ValidateMovieInput(v *Validator, movie data.Movie) *Validator {

	v.Check(NotBlank(movie.Title), "title", "must not be blank")
	v.Check(MinLength(movie.Title, 2), "title", "must be at least 2 characters long")
	v.Check(IsValidYear(movie.Year), "release_year", "must be a valid year")
	v.Check(IsValidRuntime(movie.Runtime), "runtime", "must be between 1 and 300 minutes")
	v.Check(len(movie.Genre) > 0, "genre", "must contain at least one genre")
	v.Check(AllValidGenres(movie.Genre), "genre", "contains invalid genres")
	v.Check(NotDuplicateGenre(movie.Genre), "genre", "must not contain duplicate genres")
	v.Check(AllNotBlank(movie.Actors), "actors", "actors list contains blank entries")
	v.Check(IsValidLanguage(movie.Language), "language", "invalid language")
	v.Check(IsValidCountry(movie.Country), "country", "invalid country")

	return v
}
