package http

import (
	"net/http"

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
    groupComboList, err := h.habitGroupsWithUserNames()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.HabitGroupsView(groupComboList).Render(r.Context(), w)
}

func (h *Handler) HandlePostHabitGroupView(w http.ResponseWriter, r *http.Request) {
    _, status, err := h.createFromDTO(r)
    if err != nil {
        WriteJSON(w, status, err)
        return
    }
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
