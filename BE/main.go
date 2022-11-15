package main

import (
	"database/sql"
	"db/dao"
	"db/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	// ①-2
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	db = _db
}

func accounthandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func userhandler(w http.ResponseWriter, r *http.Request) {
	CORSSetter(w)
	dao.DaoClass(w, r, db)
}

//func pointhandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Headers", "*")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//	w.Header().Set("Content-Type", "application/json")
//}

func main() {
	//http.HandleFunc("/main", messagehandler)
	http.HandleFunc("/user", userhandler)
	//http.HandleFunc("/point", pointhandler)

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

//command for getting into the database: docker exec -it db bash  ->  mysql -utest_user -ppassword test_database
