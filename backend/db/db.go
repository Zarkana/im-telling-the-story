package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	// we're not importing anything because the package doesn't want us to.
	_ "github.com/mattn/go-sqlite3"
)

// this is our location of our database. We can change this if we want, so I made it a const
const dbName string = "./db/main.db"

// Test is a test function, exported so we can call it from main
func Test() {
	// so edgy 666
	fmt.Println(NewStory(666))
}

// Takes a path name, returns true if the file exists, false otherwise
func exists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// GetConnection returns a database connection to our db
func GetConnection() *sql.DB {
	// the fact that we keep on opening db connections might be a bad idea
	// this checks if the db exists or not so we can make the proper tables
	if !exists(dbName) {
		fmt.Println("Making new Database")
		initializeDB()
	}
	fmt.Println("Opening DB")
	database, err := sql.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err)
	}
	return database
}

// this sets up our database and runs our initialization script.
func initializeDB() {
	// open the database
	database, err := sql.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err)
	}
	// this is our init SQL script
	// we might want to just slap this in here as a const, that might make it a bit more portable
	file, err := ioutil.ReadFile("./sql/init.sql")

	if err != nil {
		fmt.Println("Could not read initialization SQL script. " + err.Error())
		os.Exit(-1)
	}

	// split up the file into our different requests
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		// run every one
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

// NewStory takes the length for a story to be and returns the storyID inserted
// This does assume that the story being created is going to be the new current story.
func NewStory(length int) int64 {
	// make our db connection
	db := GetConnection()
	// be sure to close it!
	defer db.Close()
	// prepare a statement with all default values other than max length
	statement, err := db.Prepare("INSERT INTO Stories (MaxLength, StoryComplete, CurrentLength, CurrentStory) VALUES (?, false, 0, true)")
	if err != nil {
		fmt.Println(err)
	}
	// close the statement
	defer statement.Close()
	res, err := statement.Exec(length)
	if err != nil {
		fmt.Println(err)
	}
	// get the one we inserted. This might not return the right value if we're inserting a lot of stories?
	// for our use it should be fine.
	lid, _ := res.LastInsertId()

	return lid
}

// NewRound creates a new Round and returns the ID
func NewRound(storyID int, roundNum int, endTime time.Time, voteTime time.Time) {
	// make our db connection
	db := GetConnection()
	// be sure to close it!
	defer db.Close()

	// Row := db.Query()

	statement, err := db.Prepare("INSERT INTO TheRoundTable (StoryID, RoundNum, EndTime, VoteTime) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()
}
