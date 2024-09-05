package http

import (
    "net/http"
)

func (h *Handler) HandleHealthcheck(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Healthy!"))
}
