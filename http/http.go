package http

import (
    "net/http"
    "log"

    "encoding/json"
    "habits/habittrackerapi"
)

type Handler struct {
    UserService habittrackerapi.UserService
}
func NewHandler(userService habittrackerapi.UserService) *Handler {
    return &Handler{UserService: userService}
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
    router.HandleFunc("GET /api/v1/", handler.HandleHealthcheck)
    router.HandleFunc("GET /api/v1/users/{id}", handler.HandleGetUser)
    router.HandleFunc("POST /api/v1/users/", handler.HandleCreateUser)

    server := http.Server{
        Addr: s.addr,
        Handler: router,
    }
    log.Println("Listening on ", s.addr)
    return server.ListenAndServe()
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(v)
}
