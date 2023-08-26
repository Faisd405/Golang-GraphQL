package models

import (
	"database/sql"
	database "graphql-template/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Authenticate() bool {
	statement, err := database.Db.Prepare("SELECT password FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

func (user *User) CreateUser() {
	log.Println(user.Name, user.Email, user.Username, user.Password)
	statement, err := database.Db.Prepare("INSERT INTO users (name, email, username, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	hashedPadword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(user.Name, user.Email, user.Username, hashedPadword)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("SELECT id FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var id int
	err = row.Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return id, nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
