package repository

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func CreatePost(content string, createdUserID int) (int, error) {
	// defines SQL
	query := "INSERT INTO posts (content, user_id) VALUES (?, ?)"

	result, err := db.Exec(query, content, createdUserID)

	if err != nil {
		log.Printf("投稿の新規登録に失敗しました。:%v", err)
		return 0, fmt.Errorf("投稿の新規登録に失敗しました。:%w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("新規登録IDの取得に失敗しました:%v", err)
		return 0, fmt.Errorf("新規登録IDの取得二失敗しました: %w", err)
	}

	return int(id), nil
}