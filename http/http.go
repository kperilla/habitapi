package http

import (
    "net/http"
    "log"
)

type Handler struct {
}
type APIServer struct {
    addr string
    // db
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
