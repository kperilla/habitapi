package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kperilla/habitapi/habitapi"
	"github.com/kperilla/habitapi/views/templates"
)

func (h *Handler) HandleIndexView(w http.ResponseWriter, r *http.Request) {
    templates.IndexView().Render(r.Context(), w)
}

func (h *Handler) HandleGetHabitsView(w http.ResponseWriter, r *http.Request) {
    habits, err := h.HabitService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.HabitsView(habits).Render(r.Context(), w)
}

func (h *Handler) habitGroupsWithUserNames() ([]habitapi.HgUserCombo, error) {
    groups, err := h.HabitGroupService.List()
    comboList := make([]habitapi.HgUserCombo, len(groups))
    if err != nil {
        return comboList, err
    }
    for i, group := range groups {
        user, err := h.UserService.GetById(group.UserId.Hex())
        var userName string
        if err != nil {
            userName = "(error)"
        } else {
            userName = user.Name
        }
        comboList[i] = habitapi.HgUserCombo{HabitGroup: group, UserName: userName}
    }
    return comboList, err
}

func (h *Handler) HandleGetHabitGroupsView(w http.ResponseWriter, r *http.Request) {
    // groups, err := h.HabitGroupService.List()
    groupComboList, err := h.habitGroupsWithUserNames()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.HabitGroupsView(groupComboList).Render(r.Context(), w)
}

func (h *Handler) HandlePostHabitGroupView(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.CreateHabitGroupDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // TODO: Validate DTO
    validate := validator.New(validator.WithRequiredStructEnabled())
    err := validate.Struct(dto)
    if err != nil {
        fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    _, err = h.HabitGroupService.Create(dto)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // groups, err := h.HabitGroupService.List()
    groupComboList, err := h.habitGroupsWithUserNames()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.HabitGroupFormList(groupComboList).Render(r.Context(), w)
}

func (h *Handler) HandleDeleteHabitGroupView(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    err := h.HabitGroupService.Delete(id)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleGetDeedsView(w http.ResponseWriter, r *http.Request) {
    deeds, err := h.DeedService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.DeedsView(deeds).Render(r.Context(), w)
}

func (h *Handler) HandleGetRewardsView(w http.ResponseWriter, r *http.Request) {
    rewards, err := h.RewardService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.RewardsView(rewards).Render(r.Context(), w)
}
