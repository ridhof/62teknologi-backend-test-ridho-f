package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	var connectionString string = fmt.Sprintf(
		"postgresql://%v:%v@%v/%v?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASENAME"),
	)

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("fail to connect to database: %v", err)
	}

	var result int
	rows := db.QueryRow("SELECT 200 as result")
	if err := rows.Scan(
		&result,
	); err != nil || result != 200 {
			return nil, fmt.Errorf(
				"fail to validate test query %v: %v", 
				result, 
				err,
			)
	}

	fmt.Println("database connected")
	return db, nil
}

func GetInstance() (*sql.DB) {
	return db
}
