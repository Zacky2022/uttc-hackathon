package model

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func UserModel() *sql.DB {
	err := godotenv.Load("mysql.env")
	if err != nil {
		panic("Error loading .env file")
	}
	//mysqlUser := os.Getenv("MYSQL_USER")
	//mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
	//mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	//mysqlUser := os.Getenv("MYSQL_USER")
	//mysqlPwd := os.Getenv("MYSQL_PWD")
	//mysqlHost := os.Getenv("MYSQL_HOST")
	//mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	mysqlUser := "zacky"
	mysqlPwd := "HarutoMiya3/22"
	mysqlHost := "unix(/cloudsql/term2-haruto-miyazaki:us-central1:uttc)"
	mysqlDatabase := "hackathon"

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	db, err := sql.Open("mysql", connStr)

	// ①-2
	//_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase))
	//if err != nil {
	//	log.Fatalf("fail: sql.Open, %v\n", err)
	//}
	// ①-3
	//if err := _db.Ping(); err != nil {
	//	log.Fatalf("fail: _db.Ping, %v\n", err)
	//}
	//db = _db
	return db
}

func CloseOperation(db *sql.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db-kaizen.Close()")
		os.Exit(0)
	}()
}
