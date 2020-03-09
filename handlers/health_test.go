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

	t.Run("it returns 200 status code", func(t *testing.T) {
		got := response.Result().StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("it returns correct body", func(t *testing.T) {
		got := response.Body.String()
		want := `{"alive": true}`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
