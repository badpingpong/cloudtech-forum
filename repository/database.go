package repository

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func InitDB(user, password, host, port, dbname string) (err error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True", user, password, host, port, dbname,
	)

	// Connect to MySQL
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("データベース接続エラー: %v", err)
		return err
	}

	// returns DB connection information
	return db.Ping()
}

func CloseDB(){
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Printf("データベース接続解除エラー: %v", err)
		}
	}
}