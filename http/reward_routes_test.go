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

func TestHandleGetReward_ReturnsRewardMatchingId_WhenIdExists(t *testing.T) {
    var mockRewardService mock.RewardService
    var handler Handler
    handler.RewardService = &mockRewardService

    mockRewardService.GetByIdFn = func(id string) (*habitapi.Reward, error) {
        return &habitapi.Reward{Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/rewards/1", nil)

    handler.HandleGetReward(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var reward habitapi.Reward
    json.Unmarshal(w.Body.Bytes(), &reward)
    if reward.Name != "foobar" {
        t.Errorf("Expected reward name %s, got %s", "foobar", reward.Name)
    }
}

func TestHandleGetReward_ReturnsError_WhenIdDoesNotExist(t *testing.T) {
    var mockRewardService mock.RewardService
    var handler Handler
    handler.RewardService = &mockRewardService

    mockRewardService.GetByIdFn = func(id string) (*habitapi.Reward, error) {
        return nil, &habitapi.ErrResourceNotFound{}
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/rewards/1", nil)

    handler.HandleGetReward(w, r)
    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
    }
}

func TestHandleCreateReward_ReturnsId_WhenRewardCreated(t *testing.T) {
    var mockRewardService mock.RewardService
    var handler Handler
    handler.RewardService = &mockRewardService
    expectedRewardId := "1"
    postBody := bytes.NewBuffer([]byte(`{"name": "foobar"}`))

    mockRewardService.CreateFn = func(dto habitapi.CreateRewardDTO) (*habitapi.Reward, error) {
        return &habitapi.Reward{ID: expectedRewardId, Name: "foobar"}, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("POST", "/rewards/", postBody)

    handler.HandleCreateReward(w, r)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
    }

    var id string
    json.Unmarshal(w.Body.Bytes(), &id)
    if id != expectedRewardId {
        t.Errorf("Expected id %s, got %s", expectedRewardId, id)
    }
}

func TestHandleGetRewards_ReturnsAllRewardsFound_IfAnyExist(t *testing.T) {
    var mockRewardService mock.RewardService
    var handler Handler
    handler.RewardService = &mockRewardService

    mockRewardService.ListFn = func() ([]*habitapi.Reward, error) {
        rewardList := []*habitapi.Reward {
            {Name: "foobar"},
            {Name: "barfoo"},
            {Name: "barbaz"},
        }
        return rewardList, nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/rewards", nil)

    handler.HandleGetRewards(w, r)
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var rewards []*habitapi.Reward
    json.Unmarshal(w.Body.Bytes(), &rewards)
    if len(rewards) != 3 {
        t.Errorf("Expected 3 rewards, got %d", len(rewards))
    }
}

func TestHandleDeleteReward_Return204_IfNoError(t *testing.T) {
    var mockRewardService mock.RewardService
    var handler Handler
    handler.RewardService = &mockRewardService

    mockRewardService.DeleteFn = func(id string) (error) {
        return nil
    }

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("DELETE", "/rewards/1", nil)

    handler.HandleDeleteReward(w, r)

    if w.Code != http.StatusNoContent {
        t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
    }
}
