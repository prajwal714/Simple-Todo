package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func InitDB() (dbConn *sql.DB) {
	var dsn string

	dbProvider := fatalGetString("DB_PROVIDER")
	dbhost := fatalGetString("DB_HOST")
	dbPort := fatalGetString("DB_PORT")
	dbUser := fatalGetString("DB_USER")
	dbPass := fatalGetString("DB_PASSWORD")
	dbName := fatalGetString("DB_NAME")

	if dbProvider == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbPort, dbUser, dbPass, dbName)

	}
	log.Println(dsn)
	dbConn, err := sql.Open(dbUser, dsn)
	if err != nil {
		log.Error(err)
	}

	dbConn.SetMaxIdleConns(0)

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return dbConn
}
