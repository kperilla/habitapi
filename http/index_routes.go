package http

import (
	"net/http"
    "html/template"
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
