package controller

import (
	"database/sql"
	"db/usecase"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func GetController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	bytes := usecase.GetCase(db, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
