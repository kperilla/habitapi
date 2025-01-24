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

func TestHandleGetHabit_ReturnsHabitMatchingId_WhenIdExists(t *testing.T) {
    var mockHabitService mock.HabitService
    var handler Handler
    handler.HabitService = &mockHabitService

    mockHabitService.GetByIdFn = func(id string) (*habitapi.Habit, error) {
        return &habitapi.Habit{Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/habits/1", nil)

    handler.HandleGetHabit(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var habit habitapi.Habit
    json.Unmarshal(w.Body.Bytes(), &habit)
    if habit.Name != "foobar" {
        t.Errorf("Expected habit name %s, got %s", "foobar", habit.Name)
    }
}

func TestHandleGetHabit_ReturnsError_WhenIdDoesNotExist(t *testing.T) {
    var mockHabitService mock.HabitService
    var handler Handler
    handler.HabitService = &mockHabitService

    mockHabitService.GetByIdFn = func(id string) (*habitapi.Habit, error) {
        return nil, &habitapi.ErrResourceNotFound{}
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/habits/1", nil)

    handler.HandleGetHabit(w, r)
    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
    }
}

func TestHandleCreateHabit_ReturnsId_WhenHabitCreated(t *testing.T) {
    var mockHabitService mock.HabitService
    var handler Handler
    handler.HabitService = &mockHabitService
    expectedHabitId := "1"
    postBody := bytes.NewBuffer([]byte(`{"name": "foobar"}`))

    mockHabitService.CreateFn = func(dto habitapi.CreateHabitDTO) (*habitapi.Habit, error) {
        return &habitapi.Habit{ID: expectedHabitId, Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("POST", "/habits/", postBody)

    handler.HandleCreateHabit(w, r)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
    }

    var id string
    json.Unmarshal(w.Body.Bytes(), &id)
    if id != expectedHabitId {
        t.Errorf("Expected id %s, got %s", expectedHabitId, id)
    }
}

func TestHandleGetHabits_ReturnsAllHabitsFound_IfAnyExist(t *testing.T) {
    var mockHabitService mock.HabitService
    var handler Handler
    handler.HabitService = &mockHabitService

    mockHabitService.ListFn = func() ([]*habitapi.Habit, error) {
        habitList := []*habitapi.Habit {
            {Name: "foobar"},
            {Name: "barfoo"},
            {Name: "barbaz"},
        }
        return habitList, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/habits", nil)

    handler.HandleGetHabits(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var habits []*habitapi.Habit
    json.Unmarshal(w.Body.Bytes(), &habits)
    if len(habits) != 3 {
        t.Errorf("Expected 3 habits, got %d", len(habits))
    }
}

func TestHandleDeleteHabit_Return204_IfNoError(t *testing.T) {
    var mockHabitService mock.HabitService
    var handler Handler
    handler.HabitService = &mockHabitService

    mockHabitService.DeleteFn = func(id string) (error) {
        return nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("DELETE", "/habits/1", nil)

    handler.HandleDeleteHabit(w, r)

    if w.Code != http.StatusNoContent {
        t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
    }
}
