package controller

import (
	"database/sql"
	"db/usecase"
	"encoding/json"
	"github.com/oklog/ulid"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
	"unsafe"
)

type UserResForHTTPPost struct {
	Id   ulid.ULID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

func PostController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var b []byte
	b, _ = io.ReadAll(r.Body)
	var stcData usecase.StcDataType
	err := json.Unmarshal(b, &stcData)
	if err != nil {
		log.Printf("fail: could not unmarshal, %v\n", err)
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
	_ = *(*string)(unsafe.Pointer(&Id))

	usecase.PostCase(Id, db, stcData, w)
}
