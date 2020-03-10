package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)
	handler.ServeHTTP(response, request)

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
