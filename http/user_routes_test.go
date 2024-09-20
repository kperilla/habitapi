package http

import (
    "encoding/json"
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/kperilla/habitapi/habitapi"
    "github.com/kperilla/habitapi/mock"
)

func TestHandleGetUser_ReturnsUserMatchingId_WhenIdExists(t *testing.T) {
    var mockUserService mock.UserService
    var handler Handler
    handler.UserService = &mockUserService

    mockUserService.UserFn = func(id string) (*habitapi.User, error) {
        return &habitapi.User{Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/users/1", nil)

    handler.HandleGetUser(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var user habitapi.User
    json.Unmarshal(w.Body.Bytes(), &user)
    if user.Name != "foobar" {
        t.Errorf("Expected user name %s, got %s", "foobar", user.Name)
    }
}

func TestHandleGetUser_ReturnsError_WhenIdDoesNotExist(t *testing.T) {
    var mockUserService mock.UserService
    var handler Handler
    handler.UserService = &mockUserService

    mockUserService.UserFn = func(id string) (*habitapi.User, error) {
        return nil, &habitapi.ErrUserNotFound{}
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/users/1", nil)

    handler.HandleGetUser(w, r)
    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
    }
}
