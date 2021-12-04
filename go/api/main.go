package main

import (
	"log"

	"github.com/leslesnoa/go-mysql/go/api/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DB接続
	db := database.Connect()
	defer db.Close()

	// 接続確認
	err := db.Ping()
	if err != nil {
		log.Println("connection failed.")
		panic(err)
	} else {
		log.Println("connection success.")
	}

	// 行データ取得
	rows := database.GetRows(db)
	log.Println(rows)
}
