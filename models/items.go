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
	stmt, err := database.Db.Prepare("INSERT INTO items(name, description, price, image, user_id) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(item.Name, item.Description, item.Price, item.Image, item.User.ID)
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
	stmt, err := database.Db.Prepare("SELECT I.id, I.name, I.description, I.price, I.image, I.user_id, U.username FROM items I inner join users U on I.user_id = U.id")
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
	var username string
	var user_id string
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Image, &user_id, &username)
		if err != nil {
			log.Fatal(err)
		}
		item.User = &User{
			ID:       user_id,
			Username: username,
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return items
}
