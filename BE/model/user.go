package model

import (
	"database/sql"
	//"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func UserModel(db *sql.DB) *sql.DB {
	return db
}

//func UserModel() *sql.DB {
//	mysqlUser := os.Getenv("MYSQL_USER")
//	mysqlPwd := os.Getenv("MYSQL_PWD")
//	mysqlHost := os.Getenv("MYSQL_HOST")
//	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
//
//	//mysqlUser := "uttc"
//	//mysqlPwd := "HarutoMiya3/22"
//	//mysqlHost := "unix(/cloudsql/term2-haruto-miyazaki:us-central1:uttc)"
//	////mysqlHost := "34.71.244.44"
//	//mysqlDatabase := "hackathon"
//	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
//	db, err := sql.Open("mysql", connStr)
//	if err != nil {
//		log.Fatalf("fail: sql.Open, %v\n", err)
//	}
//	return db
//}

func CloseOperation(db *sql.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
