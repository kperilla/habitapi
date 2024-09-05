package main

import (
    "net/http"
)

func HandleHealthcheck(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Healthy!"))
}
