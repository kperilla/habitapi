package http

import (
	"errors"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/kperilla/habitapi/habitapi"
)

func (h *Handler) HandleCreateHabitGroup(w http.ResponseWriter, r *http.Request) {
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
    user, err := h.HabitGroupService.Create(dto)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusCreated, user.ID)
}

func (h *Handler) HandleUpdateHabitGroup(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.UpdateHabitGroupDTO
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
    user, err := h.HabitGroupService.Update(id, dto)
    if err != nil {
        // fmt.Println("UPDATE ERROR")
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, user.ID)
}

func (h *Handler) HandleGetHabitGroup(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.HabitGroupService.GetById(id)
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

func (h *Handler) HandleGetHabitGroups(w http.ResponseWriter, r *http.Request) {
    users, err := h.HabitGroupService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleDeleteHabitGroup(w http.ResponseWriter, r * http.Request) {
    id := r.PathValue("id")
    err := h.HabitGroupService.Delete(id)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, id)
}
