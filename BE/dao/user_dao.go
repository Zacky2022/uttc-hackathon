package dao

import (
	"database/sql"
	"db/controller"
	"log"
	"net/http"
)

func DaoClass(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		controller.GetController(w, db)

	case http.MethodPost:
		controller.PostController(w, r, db)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
