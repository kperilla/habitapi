package http

import (
    "net/http"
    "log"

    "encoding/json"
    "github.com/kperilla/habitapi/habitapi"
)

type Handler struct {
    UserService habitapi.UserService
}

func NewHandler(userService habitapi.UserService) *Handler {
    return &Handler{UserService: userService}
}

type APIServer struct {
    addr string
    handler *Handler
}

func NewAPIServer(addr string) *APIServer {
    return &APIServer{addr: addr}
}

func (s *APIServer) Run(handler *Handler) error {
    router := http.NewServeMux()
    router.HandleFunc("GET /", handler.HandleHealthcheck)
    router.HandleFunc("GET /users/{id}", handler.HandleGetUser)
    router.HandleFunc("GET /users/", handler.HandleGetUsers)
    router.HandleFunc("POST /users/", handler.HandleCreateUser)
    // TODO: Delete Users
    // router.HandleFunc("DELETE /users/{id}", handler.HandleDeleteUser)
    // TODO: HabitGroup
    // TODO: Habit
    // TODO: Deed
    // TODO: Reward

    v1 := http.NewServeMux()
    v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

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
