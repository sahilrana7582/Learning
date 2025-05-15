package data

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
