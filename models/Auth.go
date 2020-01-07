package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	random "github.com/alok87/goutils/pkg/random"
)

const mySQLHost = "34.93.137.151"

var mySQLConnection = fmt.Sprintf("root:password@tcp(%s)/tour_travel", mySQLHost)

//AddAuthCode function
func AddAuthCode(userID int32) (authCode string) {
	//db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tour_travel")
	db, err := sql.Open("mysql", mySQLConnection)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println(userID)
	deleteSessionQuery := `DELETE FROM session WHERE session.user_id = %d`
	deleteSessionQueryString := fmt.Sprintf(deleteSessionQuery, userID)
	fmt.Println(deleteSessionQueryString)
	// perform a db.Query delete
	delete, err := db.Query(deleteSessionQueryString)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer delete.Close()

	addSessionQuery := `INSERT INTO session(user_id, auth_code, created_at) VALUES (%d, '%s', '%s')`
	randNo := random.RangeInt(1000000000, 9999999999, 1)
	strRandNo := strconv.Itoa(randNo[0])
	t := time.Now().Format("2006-01-02 15:04:05")
	addSessionQueryString := fmt.Sprintf(addSessionQuery, userID, strRandNo, t)
	fmt.Println(addSessionQueryString)

	// perform a db.Query insert
	insert, err := db.Query(addSessionQueryString)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return strRandNo
}
