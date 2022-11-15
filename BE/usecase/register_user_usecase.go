package usecase

import (
	"database/sql"
	"encoding/json"
	"github.com/oklog/ulid"
	"log"
	"net/http"
)

type UserResForHTTPPost struct {
	Id ulid.ULID `json:"id"`
}

type StcDataType struct {
	Name string
}

func PostCase(Id ulid.ULID, db *sql.DB, stcData StcDataType, w http.ResponseWriter) {
	tx, e := db.Begin()
	if e != nil {
		log.Printf("failed to begin")
	}
	_, Error := tx.Exec("INSERT INTO user (id, name) VALUES (?,?,?)", Id.String(), stcData.Name)
	if Error != nil {
		log.Printf("fail: could not execute, %v\n", Error)
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	} else {
		var ID UserResForHTTPPost
		ID.Id = Id
		bytes, err := json.Marshal(ID)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
		tx.Commit()
		return
	}
}
