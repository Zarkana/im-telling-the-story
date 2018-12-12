package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// My linter requires a comment here
	_ "github.com/mattn/go-sqlite3"
)

const dbName string = "db/main.db"

// GetConnection returns a database connection to our db
func GetConnection() *sql.DB {
	// this checks if the db exists or not so we can make the proper tables
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		fmt.Println("Creating New DB")
		initializeDB()
	}
	database, err := sql.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err)
	}
	return database
}

func initializeDB() {
	database, err := sql.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err)
	}
	file, err := ioutil.ReadFile("./sql/init.sql")

	if err != nil {
		fmt.Println("ERROR OCCURED QUITTING " + err.Error())
		os.Exit(-1)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		result, err := database.Exec(request)
		if err != nil {
			fmt.Println(err)
			fmt.Println(result)
		}
	}
	defer database.Close()
}

// takes callsign and returns int of userid
func initSampleSubmission(name string) {

}

// takes the length for a story to be and returns the storyID
func newStory(length int) int {
	// make our db connection
	db := GetConnection()
	// be sure to close it!
	defer db.Close()
	// db.Prepare("INSERT INTO Stories ")
	// done for now i'll fix this
	return 0
}
