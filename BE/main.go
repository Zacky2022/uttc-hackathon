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
	db = model.UserModel() //&{0 0x1400012c020 0 {0 0} [] map[] 0 0 0x140001100c0 false map[] map[] 0 0 0 0 <nil> 0 0 0 0 0x1048689a0}
	//db = model.UserModel(db) //&{0 0x1400012c058 0 {0 0} [0x14000228000] map[] 0 1 0x14000110240 false map[0x14000228000:map[0x14000228000:true]] map[] 0 0 0 0 <nil> 0 0 0 0 0x1041011a0}

	log.Println(db)
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
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
