package main

import (
    "fmt"
    "net/http"
)

var counter int64 = 0

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    counter++
    fmt.Fprintf(w, "%v", counter)
}
