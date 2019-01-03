package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	Init "../sql"

	// we're not importing anything because the package doesn't want us to.
	_ "github.com/mattn/go-sqlite3"
)

// this is our location of our database. We can change this if we want, so I made it a const
const dbName string = "./main.db"

// Test is a test function, exported so we can call it from main
func Test() {
	// so edgy 666
	id := NewStory(666)
	res, _ := NewRound(id, 1, time.Now(), time.Now())
	fmt.Println(NewSubmission(res, 10))
}

// Exists Takes a file path name, returns true if the file exists, false otherwise. Exported because it might be useful
func Exists(name string) bool {
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
	if !Exists(dbName) {
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
		fmt.Println("error in opening Db")
		fmt.Println(err)
	}
	// this is our init SQL script
	// fmt.Println(Init.Init)
	// we might want to just slap this in here as a const, that might make it a bit more portable
	// file, err := ioutil.ReadFile("./sql/init.sql")

	// if err != nil {
	// 	fmt.Println("Could not read initialization SQL script. " + err.Error())
	// 	os.Exit(-1)
	// }

	// opening our golang file
	file := Init.Init

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

// NewRound creates a new Round and returns the ID or an error
func NewRound(storyID int64, roundNum int, endTime time.Time, voteTime time.Time) (int64, error) {
	// make our db connection
	db := GetConnection()
	// be sure to close it!
	defer db.Close()

	// if there already exists a round with the same story and round number, something bad has occured, so we return an error
	Row := db.QueryRow(`SELECT StoryID, RoundNum FROM TheRoundTable WHERE StoryID = ? AND RoundNum = ?`, storyID, roundNum)
	// this is kind of weird, but if there's an error here that means there are no rows, which means that there exists no such entry
	if Row.Scan() == nil {
		// you can't return nil for an int so it has to be 0?
		return 0, fmt.Errorf("there already exists a round of number %d for storyID %d", roundNum, storyID)
	}

	statement, err := db.Prepare("INSERT INTO TheRoundTable (StoryID, RoundNum, EndTime, VoteTime) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()

	res, err := statement.Exec(storyID, roundNum, endTime, voteTime)
	if err != nil {
		panic(err)
	}

	// i don't know when this would return an error
	lid, _ := res.LastInsertId()
	return lid, nil
}

// NewSubmission inserts a new submission into the database and returns the ID of the submission
func NewSubmission(roundID int64, maxLength int) int64 {
	// make our db connection
	db := GetConnection()
	// be sure to close it!
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO Submissions (Votes, Submitted, MaxLength, RoundID) VALUES (0, false, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()

	res, err := statement.Exec(maxLength, roundID)
	if err != nil {
		panic(err)
	}

	// i don't know when this would return an error
	lid, _ := res.LastInsertId()
	return lid
}

// NewUser returns the UserId of a newly inserted user
func NewUser(screenName string) int64 {
	// We should probably have functions to add the specific methods of authentication later
	// make our db connection
	db := GetConnection()
	// be sure to close it!
	defer db.Close()

	// new users shouldn't have a score higher than 0 probably?
	statement, err := db.Prepare("INSERT INTO Users (Score, ScreenName) VALUES (0, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()

	res, err := statement.Exec(screenName)
	if err != nil {
		panic(err)
	}

	// i don't know when this would return an error
	lid, _ := res.LastInsertId()
	return lid
}

// NewObjective inserts a new objective into the database and returns the ID of the objective
func NewObjective(submissionID int64, pointValue int, objectiveType string) int64 {
	// make our db connection
	db := GetConnection()
	// be sure to close it!
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO Objectives (SubmissionID, PointValue, ObjectiveType, ObjectiveMet) VALUES (?, ?, ?, false)")
	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()

	res, err := statement.Exec(submissionID, pointValue, objectiveType)
	if err != nil {
		panic(err)
	}

	// i don't know when this would return an error
	lid, _ := res.LastInsertId()
	return lid
}
