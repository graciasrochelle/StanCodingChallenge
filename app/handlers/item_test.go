package handler_test

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPostItem(t *testing.T) {
	// w := httptest.NewRecorder()
	// r := mux.NewRouter()

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode("request.json")
	
	// req, _ := http.NewRequest("POST", "/item", b)

	// handler.PostItems("Hello").AddRoute(r)
	// r.ServeHTTP(w, httptest.NewRequest("GET", "/greet/Hodor", nil))

	// if w.Code != http.StatusOK {
	// 	t.Error("Did not get expected HTTP status code, got", w.Code)
	// }
	// if w.Body.String() != "Hello Hodor!" {
	// 	t.Error("Did not get expected greeting, got", w.Body.String())
	// }

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, []byte("abcd"), w.Body.Bytes())
}
