package main

import (
	"fmt"

	"main.go/contracts"
)

var TodoList []contracts.Item

func seed() {

	var items = []contracts.Item{
		{
			Name:        "Code",
			ID:          1,
			Description: "Revive Golang Skill",
		},
		{
			Name:        "Hobby",
			ID:          2,
			Description: "Find one good hobby",
		},
		{
			Name:        "Hobby",
			ID:          3,
			Description: "Find one good hobby",
		},
	}

	for _, item := range items {
		AddItem(item)
	}

}

func main() {

	seed()
	DeleteItem(2)
	fmt.Println(TodoList)

}

func AddItem(newItem contracts.Item) {

	if newItem.ID <= 0 {
		fmt.Println("Item ID Missing")
		return
	}

	for _, item := range TodoList {
		if item.ID == newItem.ID {
			fmt.Println("Item with ID already exists !!")
			return
		}
	}

	TodoList = append(TodoList, newItem)
	fmt.Println("Item Successfully Added")
	return
}

func DeleteItem(itemID int64) {

	if itemID <= 0 {
		fmt.Println("Invalid Item ID !")
		return
	}

	for index, item := range TodoList {
		if item.ID == itemID {
			TodoList = append(TodoList[:index], TodoList[index+1:]...)
			fmt.Println("Item Successfully Deleted")
			return
		}
	}

	fmt.Println("Item Not Found")
	return

}
