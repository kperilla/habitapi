package http

import (
    "net/http"
    "errors"
    "html/template"

    "encoding/json"
    "github.com/kperilla/habitapi/habitapi"
)

type RewardsViewData struct {
    Rewards []*habitapi.Reward
}

func (h *Handler) HandleCreateReward(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.CreateRewardDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
    }
    // TODO: Validate DTO
    user, err := h.RewardService.Create(dto)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
    }
    WriteJSON(w, http.StatusCreated, user.ID)
}
func (h *Handler) HandleUpdateReward(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.UpdateRewardDTO
    id := r.PathValue("id")
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        // fmt.Println("DECODE ERROR")
        // fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // TODO: Validate DTO
    // if err := dto.Validate(); err != nil {
    //     WriteJSON(w, http.StatusBadRequest, err)
    //     return
    // }
    user, err := h.RewardService.Update(id, dto)
    if err != nil {
        // fmt.Println("UPDATE ERROR")
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, user.ID)
}

func (h *Handler) HandleGetReward(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.RewardService.GetById(id)
    var errNotFound *habitapi.ErrResourceNotFound
    switch {
        case errors.As(err, &errNotFound):
            WriteJSON(w, http.StatusNotFound, err)
        case err != nil:
            WriteJSON(w, http.StatusInternalServerError, err)
    }
    WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) HandleGetRewards(w http.ResponseWriter, r *http.Request) {
    users, err := h.RewardService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleDeleteReward(w http.ResponseWriter, r * http.Request) {
    id := r.PathValue("id")
    err := h.RewardService.Delete(id)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    WriteJSON(w, http.StatusNoContent, id)
}

func (h *Handler) HandleGetRewardsView(w http.ResponseWriter, r *http.Request) {
    viewPath := "views/rewards.html"
    t := template.Must(template.ParseFiles(viewPath))
    rewards, err := h.RewardService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    viewData := RewardsViewData{Rewards: rewards}
    err = t.Execute(w, viewData)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
}
