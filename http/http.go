package http

import (
    "net/http"
    "log"

    "encoding/json"
    "habits/habittrackerapi"
)

type Handler struct {
    userService habittrackerapi.UserService
}
type APIServer struct {
    addr string
    // db
    handler *Handler
}

func NewAPIServer(addr string) *APIServer {
    return &APIServer{addr: addr}
}

func (s *APIServer) Run(handler *Handler) error {
    router := http.NewServeMux()
    // subrouter := router.PathPrefix("/api/v1").Subrouter()
    router.HandleFunc("/", handler.HandleHealthcheck)

    server := http.Server{
        Addr: s.addr,
        Handler: router,
    }
    log.Println("Listening on ", s.addr)
    return server.ListenAndServe()
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(v)
}
