package http

import (
	"net/http"
    "html/template"

	"github.com/kperilla/habitapi/views/templates"
)

func (h *Handler) HandleIndexView(w http.ResponseWriter, r *http.Request) {
    viewPath := "views/templates/index.html"
    t := template.Must(template.ParseFiles(viewPath))
    viewData := ""
    err := t.Execute(w, viewData)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
}

func (h *Handler) HandleGetHabitsView(w http.ResponseWriter, r *http.Request) {
    habits, err := h.HabitService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.HabitsView(habits).Render(r.Context(), w)
}

func (h *Handler) HandleGetHabitGroupsView(w http.ResponseWriter, r *http.Request) {
    groups, err := h.HabitGroupService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    templates.HabitGroupsView(groups).Render(r.Context(), w)
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
