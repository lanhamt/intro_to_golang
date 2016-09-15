package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hi, the time is %s\n", time.Now())
    log.Println("serving", req.URL)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("serving on http://localhost:7777/")
    log.Fatal(http.ListenAndServe("localhost:7777", nil)) // HL
}
