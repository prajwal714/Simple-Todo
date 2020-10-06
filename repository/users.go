package repository

import (
	"time"

	"github.com/huandu/go-sqlbuilder"
	"main.go/config"
	"main.go/contracts"
)

func CreateNewUser(user contracts.User) (string, error) {
	db := config.InitDB()
	defer db.Close()

	var userID string
	user.CreatedAt = time.Now()
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("users")
	ib.Cols(
		"uuid",
		"email",
		"created_at",
	)
	ib.Values(
		user.UUID,
		user.Email,
		user.CreatedAt,
	)

	queryInsert, argsInsert := ib.Build()
	queryInsert += " RETURNING uuid"

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = db.QueryRow(queryInsert, argsInsert...).Scan(&userID)

	if err != nil {
		return "", err
	}

	return userID, nil

}
