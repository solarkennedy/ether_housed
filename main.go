package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "log/syslog"
    "sync"
)

type Common struct {
    lock sync.RWMutex
    state map[string]bool
}

func NewCommon() (c *Common, e error) {
    // load data
    return c, e
}

func (c *Common) Get(key string) (*bool, bool) {
    c.lock.RLock()
    defer c.lock.RUnlock()
    d, ok := c.state[key]
    return &d, ok
}

func (c *Common) Set(key string, d *bool) {
    c.lock.Lock()
    defer c.lock.Unlock()
    c.state[key] = *d
}

func setup_logging() {
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "ether_housed")
    if e == nil {
        log.SetOutput(logwriter)
    }
}

func main() {
    common, common_err := NewCommon()
    if common_err != nil { panic("Couldn't something") }
    http.HandleFunc("/", common.usage)
    http.HandleFunc("/on", common.usage)
    http.HandleFunc("/off", common.usage)
    http.HandleFunc("/state", common.handle_state)
    http.HandleFunc("/target_mac", common.usage)

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

func (c *Common) usage(res http.ResponseWriter, req *http.Request) {
    msg := "Welcome to ether_house."
    fmt.Fprintln(res, msg)
    log.Println("200: " + msg)
}

func (c *Common) state_handler(res http.ResponseWriter, req *http.Request) {
    params := req.URL.Query()
    api_key := params.Get("api_key")
    validate_key(api_key, 0)
}

func (c *Common) handle_state(res http.ResponseWriter, req *http.Request) {
    msg := "Welcome to state."
    fmt.Fprintln(res, msg)
    log.Println("200: " + msg)
}

func validate_key(api_key string, house_id int) (bool){
    return true
}

