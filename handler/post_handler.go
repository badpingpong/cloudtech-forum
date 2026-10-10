package handler

import (
	"encoding/json"
	"net/http"

	model "cloudtech-forum/model"
	"cloudtech-forum/repository"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	// リクエストのBodyからJSONデータを取得してpost構造体に格納
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
		return
	}

	id, err := repository.CreatePost(post.Content, post.UserID)
	if err != nil {
		http.Error(w, "投稿データの登録に失敗しました", http.StatusInternalServerError)
		return
	}

  response := map[string]interface{}{
    "message": "登録が成功しました",
    "id":      id,
  }

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}