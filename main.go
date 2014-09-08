package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "log/syslog"
)

func setup_logging() {
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "ether_housed")
    if e == nil {
        log.SetOutput(logwriter)
    }
}

func main() {
    http.HandleFunc("/", usage)
    http.HandleFunc("/on", usage)
    http.HandleFunc("/off", usage)
    http.HandleFunc("/state", usage)
    http.HandleFunc("/target_mac", usage)

    port := os.Getenv("PORT")
    if port == ""{
        port = "3000"
        log.Print("No PORT variable. Defaulting to 3000")
    }
    log.Print("listening on " + port + "..." )
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
      panic(err)
    }
}

func usage(res http.ResponseWriter, req *http.Request) {
    msg := "Welcome to ether_house."
    fmt.Fprintln(res, msg)
    fmt.Println("200: " + msg)
}
