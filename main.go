package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", hello)
    port := os.Getenv("PORT")
    if port == ""{
        panic("No PORT variable. Please set one")
    }
    fmt.Println("listening on " + port + "..." )
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
    fmt.Println("200: hello, world")
}
