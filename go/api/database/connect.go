package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int
	Name       string
	Created_at sql.NullString
	updated_at sql.NullString
}

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("start dbConnectFunc.")
	// user := "test"
	// password := "test"
	// host := "localhost"
	// port := "3306"
	// dbName := "testdb"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	conn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4"

	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetRows(db *sql.DB) []User {
	log.Println("start GetRowsFunc.")
	table := os.Getenv("DB_TABLE")
	cmd := "SELECT * FROM " + table
	rows, _ := db.Query(cmd)
	var result []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Created_at, &u.updated_at)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, u)
	}
	for _, u := range result {
		log.Println(u.Id, u.Name)
	}
	return result
}

func GetRowById(db *sql.DB) {

}
