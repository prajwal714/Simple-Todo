package main

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func Load(router *mux.Router, dbConn *sql.DB, dbRedis *redis.Client) {

	log.Println("Load function")
}
