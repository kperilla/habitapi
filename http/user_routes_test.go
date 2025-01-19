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

    mockUserService.CreateUserFn = func(dto habitapi.CreateUserDTO) (*habitapi.User, string, error) {
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

func TestHandleGetUsers_ReturnsAllUsersFound_IfAnyExist(t *testing.T) {
    var mockUserService mock.UserService
    var handler Handler
    handler.UserService = &mockUserService

    mockUserService.UsersFn = func() ([]*habitapi.User, error) {
        userList := []*habitapi.User {
            {Name: "foobar"},
            {Name: "barfoo"},
            {Name: "barbaz"},
        }
        return userList, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/users", nil)

    handler.HandleGetUsers(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var users []*habitapi.User
    json.Unmarshal(w.Body.Bytes(), &users)
    if len(users) != 3 {
        t.Errorf("Expected 3 users, got %d", len(users))
    }
}

func TestHandleDeleteUser_Return204_IfNoError(t *testing.T) {
    var mockUserService mock.UserService
    var handler Handler
    handler.UserService = &mockUserService

    mockUserService.DeleteUserFn = func(id string) (error) {
        return nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("DELETE", "/users/1", nil)

    handler.HandleDeleteUser(w, r)

    if w.Code != http.StatusNoContent {
        t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
    }
}
