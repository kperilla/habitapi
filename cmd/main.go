package main

import (
    "log"

    "example.com/habits/http"
)

func main() {
    server := http.NewAPIServer(":8080")
    var handler http.Handler
    if err := server.Run(&handler); err != nil {
        log.Fatal(err)
    }
}
