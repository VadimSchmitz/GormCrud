package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqliteDatabase, _ := sql.Open("sqlite3", "./movies.db")
	defer sqliteDatabase.Close()

	arguments := os.Args[1:]
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addImdbId := addCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")
	addTitle := addCommand.String("title", "Carmencita", "The movie's or series' title")
	addYear := addCommand.Int("year", 1894, "The movie's or series' year of release")
	addImdbRating := addCommand.Float64("rating", 5.7, "The movie's or series' rating on IMDb")

	detailsCommand := flag.NewFlagSet("details", flag.ExitOnError)
	detailsImdbId := detailsCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteImdbId := deleteCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")

	switch arguments[0] {
	case "add":
		addCommand.Parse(arguments[1:])
		insertMovie(sqliteDatabase, *addImdbId, *addTitle, *addImdbRating, *addYear)
	case "list":
		displayAllMovieTitles(sqliteDatabase)
	case "details":
		detailsCommand.Parse(arguments[1:])
		movieDetails(sqliteDatabase, *detailsImdbId)
	case "delete":
		deleteCommand.Parse(arguments[1:])
		deleteMovieById(sqliteDatabase, *deleteImdbId)
	}
}

func displayAllMovieTitles(db *sql.DB) {
	row, err := db.Query("SELECT Title FROM 'movies'")
	checkError(err)
	defer row.Close()

	for row.Next() {
		var title string
		row.Scan(&title)
		fmt.Println(title)
	}
}

func movieDetails(db *sql.DB, imdb_id string) {
	var Movie struct {
		Imdb_id string
		Title   string
		Rating  float64
		Year    int
	}

	row := db.QueryRow("SELECT * FROM movies WHERE IMDb_id = ?", imdb_id)
	err := row.Scan(&Movie.Imdb_id, &Movie.Title, &Movie.Rating, &Movie.Year)
	checkError(err)

	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d", Movie.Imdb_id, Movie.Title, Movie.Rating, Movie.Year)
}

func deleteMovieById(db *sql.DB, imdb_id string) {
	row, err := db.Prepare("DELETE FROM movies WHERE IMDb_id = ?")
	checkError(err)
	res, err := row.Exec(imdb_id)
	checkError(err)
	a, err := res.RowsAffected()
	checkError(err)
	if a == 1 {
		fmt.Println("Movie deleted")
	}
}

func insertMovie(db *sql.DB, imdb_id string, title string, rating float64, year int) {
	statement, err := db.Prepare("INSERT INTO movies(IMDb_id, Title, Rating, Year) VALUES(?, ?, ?, ?)")
	checkError(err)

	_, err = statement.Exec(imdb_id, title, rating, year)
	checkError(err)
	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d", imdb_id, title, rating, year)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
