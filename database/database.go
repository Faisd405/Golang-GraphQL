package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

// Define the database connection parameters.
const (
	DBUser     = "root"
	DBPassword = ""
	DBName     = "graphql-go"
	DBHost     = ""
)

func InitDB() {
	var err error
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	dataSourceName := fmt.Sprintf("%s:%s@%s/%s", DBUser, DBPassword, DBHost, DBName)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
