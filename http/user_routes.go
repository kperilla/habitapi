package http

import (
	"errors"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/kperilla/habitapi/habitapi"
)

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.CreateUserDTO
    fmt.Println("Calling Handle Create User")
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        fmt.Println("Decode Error")
        fmt.Println(err.Error())
        // fmt.Println(json.NewDecoder(r.Body))
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // TODO: Validate DTO
    if err := dto.Validate(); err != nil {
        fmt.Println(err)
        fmt.Println("Validate Error")
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    user, err := h.UserService.Create(dto)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Create Error")
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusCreated, user.ID)
}

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.UserService.GetById(id)
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

func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HandleGetUsers Called")
    users, err := h.UserService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.UpdateUserDTO
    id := r.PathValue("id")
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        fmt.Println("DECODE ERROR")
        fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // TODO: Validate DTO
    if err := dto.Validate(); err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    user, err := h.UserService.Update(id, dto)
    if err != nil {
        fmt.Println("UPDATE ERROR")
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, user.ID)
}

func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r * http.Request) {
    id := r.PathValue("id")
    err := h.UserService.Delete(id)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, id)
}
