package usecase

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetCase(db *sql.DB, w http.ResponseWriter) []byte {
	rows, err := db.Query("SELECT id,name FROM user")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	// ②-3
	users := make([]UserResForHTTPGet, 0)
	for rows.Next() {
		var u UserResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
			return nil
		}
		users = append(users, u)
	}
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	return bytes
}
