package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

type APIServer struct {
    addr string
    // db
}

func NewAPIServer(addr string) *APIServer {
    return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
    router := mux.NewRouter()
    // subrouter := router.PathPrefix("/api/v1").Subrouter()

    log.Println("Listening on ", s.addr)
    return http.ListenAndServe(s.addr, router)
}
