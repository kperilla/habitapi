package main

import (
    "log"

    "habits/http"
)

func main() {
    server := http.NewAPIServer(":8080")
    var handler http.Handler
    if err := server.Run(&handler); err != nil {
        log.Fatal(err)
    }
}
