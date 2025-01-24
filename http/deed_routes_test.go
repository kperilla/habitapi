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

func TestHandleGetDeed_ReturnsDeedMatchingId_WhenIdExists(t *testing.T) {
    var mockDeedService mock.DeedService
    var handler Handler
    handler.DeedService = &mockDeedService

    mockDeedService.GetByIdFn = func(id string) (*habitapi.Deed, error) {
        return &habitapi.Deed{Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/deeds/1", nil)

    handler.HandleGetDeed(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var deed habitapi.Deed
    json.Unmarshal(w.Body.Bytes(), &deed)
    if deed.Name != "foobar" {
        t.Errorf("Expected deed name %s, got %s", "foobar", deed.Name)
    }
}

func TestHandleGetDeed_ReturnsError_WhenIdDoesNotExist(t *testing.T) {
    var mockDeedService mock.DeedService
    var handler Handler
    handler.DeedService = &mockDeedService

    mockDeedService.GetByIdFn = func(id string) (*habitapi.Deed, error) {
        return nil, &habitapi.ErrResourceNotFound{}
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/deeds/1", nil)

    handler.HandleGetDeed(w, r)
    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
    }
}

func TestHandleCreateDeed_ReturnsId_WhenDeedCreated(t *testing.T) {
    var mockDeedService mock.DeedService
    var handler Handler
    handler.DeedService = &mockDeedService
    expectedDeedId := "1"
    postBody := bytes.NewBuffer([]byte(`{"name": "foobar"}`))

    mockDeedService.CreateFn = func(dto habitapi.CreateDeedDTO) (*habitapi.Deed, error) {
        return &habitapi.Deed{ID: expectedDeedId, Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("POST", "/deeds/", postBody)

    handler.HandleCreateDeed(w, r)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
    }

    var id string
    json.Unmarshal(w.Body.Bytes(), &id)
    if id != expectedDeedId {
        t.Errorf("Expected id %s, got %s", expectedDeedId, id)
    }
}

func TestHandleGetDeeds_ReturnsAllDeedsFound_IfAnyExist(t *testing.T) {
    var mockDeedService mock.DeedService
    var handler Handler
    handler.DeedService = &mockDeedService

    mockDeedService.ListFn = func() ([]*habitapi.Deed, error) {
        deedList := []*habitapi.Deed {
            {Name: "foobar"},
            {Name: "barfoo"},
            {Name: "barbaz"},
        }
        return deedList, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/deeds", nil)

    handler.HandleGetDeeds(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var deeds []*habitapi.Deed
    json.Unmarshal(w.Body.Bytes(), &deeds)
    if len(deeds) != 3 {
        t.Errorf("Expected 3 deeds, got %d", len(deeds))
    }
}

func TestHandleDeleteDeed_Return204_IfNoError(t *testing.T) {
    var mockDeedService mock.DeedService
    var handler Handler
    handler.DeedService = &mockDeedService

    mockDeedService.DeleteFn = func(id string) (error) {
        return nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("DELETE", "/deeds/1", nil)

    handler.HandleDeleteDeed(w, r)

    if w.Code != http.StatusNoContent {
        t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
    }
}
