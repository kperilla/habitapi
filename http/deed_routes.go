package http

import (
	"errors"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/kperilla/habitapi/habitapi"
)

type DeedsViewData struct {
    Deeds []*habitapi.Deed
}

func (h *Handler) HandleCreateDeed(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.CreateDeedDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // TODO: Validate DTO
    user, err := h.DeedService.Create(dto)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusCreated, user.ID)
}
func (h *Handler) HandleUpdateDeed(w http.ResponseWriter, r *http.Request) {
    var dto habitapi.UpdateDeedDTO
    id := r.PathValue("id")
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        fmt.Println(err)
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    // TODO: Validate DTO
    // if err := dto.Validate(); err != nil {
    //     WriteJSON(w, http.StatusBadRequest, err)
    //     return
    // }
    user, err := h.DeedService.Update(id, dto)
    if err != nil {
        // fmt.Println("UPDATE ERROR")
        WriteJSON(w, http.StatusBadRequest, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, user.ID)
}

func (h *Handler) HandleGetDeed(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.DeedService.GetById(id)
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

func (h *Handler) HandleGetDeeds(w http.ResponseWriter, r *http.Request) {
    users, err := h.DeedService.List()
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleDeleteDeed(w http.ResponseWriter, r * http.Request) {
    id := r.PathValue("id")
    err := h.DeedService.Delete(id)
    if err != nil {
        WriteJSON(w, http.StatusInternalServerError, err)
        return
    }
    WriteJSON(w, http.StatusNoContent, id)
}
