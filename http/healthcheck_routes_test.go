package http

import (
    "testing"
    "net/http"
	"net/http/httptest"
)

func TestHandleHealthcheck(t *testing.T) {
    var handler Handler
    w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
    handler.HandleHealthcheck(w, r)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }
}
