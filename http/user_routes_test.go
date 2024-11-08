package http

import (
    "bytes"
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

func TestHandleCreateUser_ReturnsId_WhenUserCreated(t *testing.T) {
    var mockUserService mock.UserService
    var handler Handler
    handler.UserService = &mockUserService
    expectedUserId := "1"
    postBody := bytes.NewBuffer([]byte(`{"name": "foobar"}`))

    mockUserService.CreateUserFn = func(name string) (*habitapi.User, string, error) {
        return &habitapi.User{Name: "foobar"}, expectedUserId, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("POST", "/users/", postBody)

    handler.HandleCreateUser(w, r)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
    }

    var id string
    json.Unmarshal(w.Body.Bytes(), &id)
    if id != expectedUserId {
        t.Errorf("Expected id %s, got %s", expectedUserId, id)
    }
}
