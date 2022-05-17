package model

type Genre struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Movie struct {
	ID           int    `db:"id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	Year         int    `db:"year"`
	GenreName    string `db:"genre_name"`
	DirectorName string `db:"director_name"`
}

type Director struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Birthday string `db:"birthday"`
}

type MovieDirector struct {
	ID           int     `db:"id"`
	Title        string  `db:"title"`
	Description  string  `db:"description"`
	Year         int     `db:"year"`
	DirectorName *string `db:"director_name"`
}
