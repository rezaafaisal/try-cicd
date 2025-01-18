package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "apa kabar!")
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "halo! wkwkwk from test doang")
    })

    http.HandleFunc("/index", index)

    fmt.Println("starting web server at http://localhost:6100/")
    http.ListenAndServe(":6100", nil)
}