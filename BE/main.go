package main

import (
	"database/sql"
	"db/dao"
	"db/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	db = model.UserModel(db)
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	dao.DaoClass(w, r, db)

}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	model.CloseOperation(db)
}
