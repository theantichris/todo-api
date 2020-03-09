package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealth(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	router.ServeHTTP(response, req)

	got := response.Body.String()
	want := `{"alive": true}`

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
