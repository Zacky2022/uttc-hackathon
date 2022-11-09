package controller_test

import (
	"bytes"
	"database/sql"
	"db/controller"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var PostBody1 = []byte(`{"name": "", "age": 34}`)
var PostBody2 = []byte(`{"name":"ichiro", "age":128}`)
var db *sql.DB

func TestPostController(t *testing.T) {
	t.Run("NameBadRequest", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/user?name=ichiro", bytes.NewReader(PostBody1))
		rec := httptest.NewRecorder()
		controller.PostController(rec, req, db)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("AgeBadRequest", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/user?name=ichiro", bytes.NewReader(PostBody2))
		rec := httptest.NewRecorder()
		controller.PostController(rec, req, db)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
