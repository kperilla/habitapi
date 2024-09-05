package main

import (
    "net/http"
    "log"
)

type APIServer struct {
    addr string
    // db
}

func NewAPIServer(addr string) *APIServer {
    return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
    router := http.NewServeMux()
    // subrouter := router.PathPrefix("/api/v1").Subrouter()
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Healthy!"))
    })

    server := http.Server{
        Addr: s.addr,
        Handler: router,
    }
    log.Println("Listening on ", s.addr)
    return server.ListenAndServe()
}
