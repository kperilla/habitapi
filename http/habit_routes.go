package http

import (
	"errors"
	"fmt"
	"net/http"
    "html/template"

	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/kperilla/habitapi/habitapi"
)

type HabitsViewData struct {
    Habits []*habitapi.Habit
}

func (h *Handler) HandleCreateHabit(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.CreateHabitDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // TODO: Give more descriptive message to client
    validate := validator.New(validator.WithRequiredStructEnabled())
    err := validate.Struct(dto)
    if err != nil {
        fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    habit, err := h.HabitService.Create(dto)
    if err != nil {
        fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusCreated, habit.ID)
}

func (h *Handler) HandleUpdateHabit(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.UpdateHabitDTO
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
    user, err := h.HabitService.Update(id, dto)
    if err != nil {
        // fmt.Println("UPDATE ERROR")
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, user.ID)
}

func (h *Handler) HandleGetHabit(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.HabitService.GetById(id)
    var errNotFound *habitapi.ErrResourceNotFound
    switch {
        case errors.As(err, &errNotFound):
            WriteJSON(w, http.StatusNotFound, err)
            return
        case err != nil:
            WriteJSON(w, http.StatusInternalServerError, err)
            return
    }
    WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) HandleGetHabits(w http.ResponseWriter, r *http.Request) {
    users, err := h.HabitService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleDeleteHabit(w http.ResponseWriter, r * http.Request) {
    id := r.PathValue("id")
    err := h.HabitService.Delete(id)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, id)
}

func (h *Handler) HandleGetHabitsView(w http.ResponseWriter, r *http.Request) {
    viewPath := "views/habits.html"
    t := template.Must(template.ParseFiles(viewPath))
    habits, err := h.HabitService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    viewData := HabitsViewData{Habits: habits}
    err = t.Execute(w, viewData)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
}
