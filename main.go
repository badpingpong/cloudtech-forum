package main
import (
	"log"
	"net/http"
	"os"

	"cloudtech-forum/handler"
	"cloudtech-forum/repository"

	"github.com/gorilla/mux"
  	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	apiport := os.Getenv("API_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD") // DBのパスワード
	host := os.Getenv("DB_HOST")         // DBのホスト
	port := os.Getenv("DB_PORT")         // DBのポート
	dbname := os.Getenv("DB_NAME")       // DB名

	err := repository.InitDB(username, password, host, port, dbname)
	if err != nil {
    	log.Fatalf("データベースに接続できません: %v", err)
  	}
  	defer repository.CloseDB() // プログラム終了時にデータベース接続を閉じる

	r := mux.NewRouter()
	r.HandleFunc("/posts", handler.CreateHandler).Methods("POST")

	log.Println("APIサーバを起動しました。ポート: " + apiport)
	if err := http.ListenAndServe(":"+apiport, r); err != nil {
    	log.Fatal(err)
	}
}