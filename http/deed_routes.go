package http

import (
    "net/http"
    "errors"

    "encoding/json"
    "github.com/kperilla/habitapi/habitapi"
)

func (h *Handler) HandleCreateDeed(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.CreateDeedDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
    }
    // TODO: Validate DTO
    user, err := h.DeedService.Create(dto)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
    }
    WriteJSON(w, http.StatusCreated, user.ID)
}

func (h *Handler) HandleGetDeed(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.DeedService.GetById(id)
    var errNotFound *habitapi.ErrResourceNotFound
    switch {
        case errors.As(err, &errNotFound):
            WriteJSON(w, http.StatusNotFound, err)
        case err != nil:
            WriteJSON(w, http.StatusInternalServerError, err)
    }
    WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) HandleGetDeeds(w http.ResponseWriter, r *http.Request) {
    users, err := h.DeedService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleDeleteDeed(w http.ResponseWriter, r * http.Request) {
    id := r.PathValue("id")
    err := h.DeedService.Delete(id)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
    }
    WriteJSON(w, http.StatusNoContent, id)
}
