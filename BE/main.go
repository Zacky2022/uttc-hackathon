package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type UserResForHTTPGet struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Point int    `json:"point"`
}

type UserResForHTTPPost struct {
	Res string `json:"id"`
}

type stcDataType struct {
	Name  string
	Point int
}

type ConResForGet struct {
	Msid    string `json:"msid"`
	Point   int    `json:"sentpoint"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

type ConsPOST struct {
	From    string
	To      string
	Point   int
	Message string
}

type Updatetype struct {
	Targ    string `json:"targ"`
	Point   int    `json:"point"`
	Message string `json:"message"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	db = _db
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func userhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "https://uttc-hackathon-kappa.vercel.app/")
	w.Header().Set("Access-Control-Allow-Origin", "https://uttc-hackathon-kappa.vercel.app/")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT id, name,point FROM user")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ②-3
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Point); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// ②-4
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	case http.MethodPost:
		var stcData stcDataType
		err := json.NewDecoder(r.Body).Decode(&stcData)
		if err != nil {
			log.Printf("fail: could not decode, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if stcData.Name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		Id := ulid.MustNew(ulid.Timestamp(t), entropy)
		tx, err := db.Begin()
		if err != nil {
			log.Printf("failed to begin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, Error := tx.Exec("INSERT INTO user (id, name, point) VALUES (?,?,0)", Id.String(), stcData.Name)
		if Error != nil {
			log.Printf("fail: could not execute, %v\n", Error)
			w.WriteHeader(http.StatusInternalServerError)
			err := tx.Rollback()
			if err != nil {
				log.Println("failed to Rollback")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		var ID UserResForHTTPPost
		ID.Res = Id.String()
		bytes, err := json.Marshal(ID)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
		if Err := tx.Commit(); Err != nil {
			log.Println("failed to commit")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tx.Commit()
		return

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func listhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "https://uttc-hackathon-kappa.vercel.app/")
	w.Header().Set("Access-Control-Allow-Origin", "https://uttc-hackathon-kappa.vercel.app/")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		ft := r.URL.Query().Get("ft")
		if ft != "from" && ft != "to" {
			log.Println("fail: invalid method")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if ft == "from" {
			rows, err := db.Query("SELECT msid, sentpoint, message, name FROM contribution JOIN user ON idto=id WHERE idfrom=?", userId)
			if err != nil {
				log.Printf("fail: db.Query, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users := make([]ConResForGet, 0)
			for rows.Next() {
				var u ConResForGet
				if err := rows.Scan(&u.Msid, &u.Point, &u.Message, &u.Name); err != nil {
					log.Printf("fail: rows.Scan, %v\n", err)

					if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
						log.Printf("fail: rows.Close(), %v\n", err)
					}
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				users = append(users, u)
			}
			bytes, err := json.Marshal(users)
			if err != nil {
				log.Printf("fail: json.Marshal, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)
		} else if ft == "to" {
			rows, err := db.Query("SELECT sentpoint, message, name FROM contribution JOIN user ON idfrom=id WHERE idto=?", userId)
			if err != nil {
				log.Printf("fail: db.Query, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users := make([]ConResForGet, 0)
			for rows.Next() {
				var u ConResForGet
				if err := rows.Scan(&u.Point, &u.Message, &u.Name); err != nil {
					log.Printf("fail: rows.Scan, %v\n", err)

					if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
						log.Printf("fail: rows.Close(), %v\n", err)
					}
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				users = append(users, u)
			}
			bytes, err := json.Marshal(users)
			if err != nil {
				log.Printf("fail: json.Marshal, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)
		} else {
			log.Println("fail: invalid query parameter")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var consPost ConsPOST
		err := json.NewDecoder(r.Body).Decode(&consPost)
		if err != nil {
			log.Printf("fail: could not decode, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if consPost.To == "" {
			log.Println("fail: No address defined")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		Id := ulid.MustNew(ulid.Timestamp(t), entropy)
		tx, err := db.Begin()
		if err != nil {
			log.Printf("failed to begin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, Error := tx.Exec("INSERT INTO contribution VALUES (?,?,?,?,?)", Id.String(), userId, consPost.To, consPost.Point, consPost.Message)
		if Error != nil {
			log.Printf("fail: could not execute, %v\n", Error)
			w.WriteHeader(http.StatusInternalServerError)
			err := tx.Rollback()
			if err != nil {
				log.Println("failed to Rollback")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		_, Err := tx.Exec("UPDATE user SET point=(SELECT SUM(sentpoint) FROM contribution WHERE idto=?) WHERE id=?", consPost.To, consPost.To)
		if Err != nil {
			log.Println("failed to update point")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ID UserResForHTTPPost
		ID.Res = Id.String()
		bytes, err := json.Marshal(ID)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
		if Err := tx.Commit(); Err != nil {
			log.Println("failed to commit")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tx.Commit()
		return

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func updatehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "https://uttc-hackathon-kappa.vercel.app/")
	w.Header().Set("Access-Control-Allow-Origin", "https://uttc-hackathon-kappa.vercel.app/")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT msid, sentpoint, message, name FROM contribution JOIN user ON idto=id")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users := make([]Updatetype, 0)
		for rows.Next() {
			var u Updatetype
			if err := rows.Scan(&u.Targ, &u.Point, &u.Message); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	case http.MethodPost:
		var updatebody Updatetype
		err := json.NewDecoder(r.Body).Decode(&updatebody)
		if err != nil {
			log.Printf("fail: could not decode, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if updatebody.Targ == "" {
			log.Println("fail: No address defined")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		Id := ulid.MustNew(ulid.Timestamp(t), entropy)
		tx, err := db.Begin()
		if err != nil {
			log.Printf("failed to begin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, Error := tx.Exec("UPDATE contribution SET sentpoint=?,message=? WHERE msid=?", updatebody.Point, updatebody.Message, updatebody.Targ)
		if Error != nil {
			log.Printf("fail: could not execute, %v\n", Error)
			w.WriteHeader(http.StatusInternalServerError)
			err := tx.Rollback()
			if err != nil {
				log.Println("failed to Rollback")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		_, Err := tx.Exec("UPDATE user SET point=(SELECT SUM(sentpoint) FROM contribution WHERE idto=(SELECT idto FROM contribution WHERE msid=?)) WHERE id=(SELECT idto FROM contribution WHERE msid=?)", updatebody.Targ, updatebody.Targ)
		if Err != nil {
			log.Println("failed to update point")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ID UserResForHTTPPost
		ID.Res = Id.String()
		bytes, err := json.Marshal(ID)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
		if Err := tx.Commit(); Err != nil {
			log.Println("failed to commit")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tx.Commit()
		return

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", userhandler)
	http.HandleFunc("/con-list", listhandler)
	http.HandleFunc("/update", updatehandler)

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
