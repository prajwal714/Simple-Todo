package repository

import (
	"github.com/huandu/go-sqlbuilder"
	"main.go/config"
	"main.go/contracts"
)

func Insert(item contracts.Item) (int64, error) {

	db := config.InitDB()
	defer db.Close()
	var itemID int64

	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("items")
	ib.Cols(

		"name",
		"description",
	)

	ib.Values(

		item.Name,
		item.Description,
	)

	queryInsert, argsInsert := ib.Build()
	queryInsert += " RETURNING id"
	//log.Println(queryInsert)
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	err = db.QueryRow(queryInsert, argsInsert...).Scan(&itemID)

	if err != nil {
		return -1, err
	}

	return itemID, nil

}

func List() ([]*contracts.Item, error) {
	db := config.InitDB()
	defer db.Close()

	result := make([]*contracts.Item, 0)

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "name", "description")
	sb.From("items")

	query, args := sb.Build()

	rows, err := db.Query(query, args...)
	if err != nil {
		return result, nil
	}

	for rows.Next() {
		item := new(contracts.Item)

		err = rows.Scan(&item.ID, &item.Name, &item.Description)

		result = append(result, item)

	}

	return result, nil

}

func Delete(id int64) error {
	db := config.InitDB()
	defer db.Close()

	delb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	delb.DeleteFrom("items")
	delb.Where(delb.Equal("id", id))

	query, args := delb.Build()

	_, err := db.Query(query, args...)

	if err != nil {
		return err
	}

	return nil

}
