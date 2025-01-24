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

func TestHandleGetHabitGroup_ReturnsHabitGroupMatchingId_WhenIdExists(t *testing.T) {
    var mockHabitGroupService mock.HabitGroupService
    var handler Handler
    handler.HabitGroupService = &mockHabitGroupService

    mockHabitGroupService.GetByIdFn = func(id string) (*habitapi.HabitGroup, error) {
        return &habitapi.HabitGroup{Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/habit_groups/1", nil)

    handler.HandleGetHabitGroup(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var habitGroup habitapi.HabitGroup
    json.Unmarshal(w.Body.Bytes(), &habitGroup)
    if habitGroup.Name != "foobar" {
        t.Errorf("Expected habitGroup name %s, got %s", "foobar", habitGroup.Name)
    }
}

func TestHandleGetHabitGroup_ReturnsError_WhenIdDoesNotExist(t *testing.T) {
    var mockHabitGroupService mock.HabitGroupService
    var handler Handler
    handler.HabitGroupService = &mockHabitGroupService

    mockHabitGroupService.GetByIdFn = func(id string) (*habitapi.HabitGroup, error) {
        return nil, &habitapi.ErrResourceNotFound{}
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/habit_groups/1", nil)

    handler.HandleGetHabitGroup(w, r)
    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
    }
}

func TestHandleCreateHabitGroup_ReturnsId_WhenHabitGroupCreated(t *testing.T) {
    var mockHabitGroupService mock.HabitGroupService
    var handler Handler
    handler.HabitGroupService = &mockHabitGroupService
    expectedHabitGroupId := "1"
    postBody := bytes.NewBuffer([]byte(`{"name": "foobar"}`))

    mockHabitGroupService.CreateFn = func(dto habitapi.CreateHabitGroupDTO) (*habitapi.HabitGroup, error) {
        return &habitapi.HabitGroup{ID: expectedHabitGroupId, Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("POST", "/habit_groups/", postBody)

    handler.HandleCreateHabitGroup(w, r)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
    }

    var id string
    json.Unmarshal(w.Body.Bytes(), &id)
    if id != expectedHabitGroupId {
        t.Errorf("Expected id %s, got %s", expectedHabitGroupId, id)
    }
}

func TestHandleGetHabitGroups_ReturnsAllHabitGroupsFound_IfAnyExist(t *testing.T) {
    var mockHabitGroupService mock.HabitGroupService
    var handler Handler
    handler.HabitGroupService = &mockHabitGroupService

    mockHabitGroupService.ListFn = func() ([]*habitapi.HabitGroup, error) {
        habitGroupList := []*habitapi.HabitGroup {
            {Name: "foobar"},
            {Name: "barfoo"},
            {Name: "barbaz"},
        }
        return habitGroupList, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/habit_groups", nil)

    handler.HandleGetHabitGroups(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var habitGroups []*habitapi.HabitGroup
    json.Unmarshal(w.Body.Bytes(), &habitGroups)
    if len(habitGroups) != 3 {
        t.Errorf("Expected 3 habitGroups, got %d", len(habitGroups))
    }
}

func TestHandleDeleteHabitGroup_Return204_IfNoError(t *testing.T) {
    var mockHabitGroupService mock.HabitGroupService
    var handler Handler
    handler.HabitGroupService = &mockHabitGroupService

    mockHabitGroupService.DeleteFn = func(id string) (error) {
        return nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("DELETE", "/habit_groups/1", nil)

    handler.HandleDeleteHabitGroup(w, r)

    if w.Code != http.StatusNoContent {
        t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
    }
}
