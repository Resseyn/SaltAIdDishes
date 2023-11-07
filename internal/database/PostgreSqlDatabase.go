package database

import (
	"SaltAIdDishes/pkg/loggers"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var GlobalDatabase *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "testtest"
	dbname   = "dishesDatabase"
)

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	GlobalDatabase, err = openDB(dsn)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
