package controller

import (
	"database/sql"
	"db/usecase"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func GetController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	name := r.URL.Query().Get("name")
	if name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes := usecase.GetCase(db, w, name)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
