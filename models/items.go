package models

import (
	database "graphql-template/database"
	"log"
)

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	User        *User   `json:"user"`
}

// #2
func (item Item) SaveItem() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO items(name, description, price, image) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(item.Name, item.Description, item.Price, item.Image)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAllItems() []Item {
	stmt, err := database.Db.Prepare("SELECT id, name, description, price, image FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Image)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return items
}
