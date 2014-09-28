package main

import (
	"fmt"
	"log"
	"log/syslog"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const NUM_HOUSES = 8

type common struct {
	lock       sync.RWMutex
	state      []bool
	api_key    []string
	target_mac []string
}

var Common = new(common)

/*
func NewCommon() (c *Common, e error) {
	cstate := load_existing_state()
	fmt.Println(cstate[0])
//	copy(cstate, c.state)
	c.state = cstate
	return c, e
}
*/

func load_existing_state() {
	// TODO: Get state out of memcache
	//        Common.state = []bool {true, true, true, true, false, false, false, false}
	Common.state = []bool{false, false, false, false, true, true, true, true}
	// Common.state = []bool {true, false, true, false, true, false, true, false}
	return
}

func load_secrets() {
	Common.api_key = []string{"", "", "", "", "", "", "", ""}
	for i := 0; i < NUM_HOUSES; i++ {
		Common.api_key[i] = os.Getenv("APIKEY" + strconv.Itoa(i))
		if Common.api_key[i] == "" {
			log.Println("WARNING: Didn't get an API key for " + strconv.Itoa(i) + ".")
		}
		log.Println("INFO: API Key for " + strconv.Itoa(i) + " is " + Common.api_key[i])
	}
}

func Get(id int) *bool {
	Common.lock.RLock()
	defer Common.lock.RUnlock()
	d := Common.state[id]
	return &d
}

func Set(id int, d *bool) {
	Common.lock.Lock()
	defer Common.lock.Unlock()
	Common.state[id] = *d
}

func setup_logging() {
	logwriter, e := syslog.New(syslog.LOG_NOTICE, "ether_housed")
	if e == nil {
		log.SetOutput(logwriter)
	}
}

func main() {
	load_existing_state()
	load_secrets()
	http.HandleFunc("/", usage)
	http.HandleFunc("/on", turn_on)
	http.HandleFunc("/off", turn_off)
	http.HandleFunc("/state", handle_state)
	http.HandleFunc("/target_mac", target_mac_handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Print("No PORT variable. Defaulting to 3000")
	}
	log.Print("listening on " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func usage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		log.Println("404: " + req.URL.Path)
		return
	}
	msg := "Welcome to ether_house."
	fmt.Fprintln(res, msg)
	log.Println("200: " + req.URL.Path)
}

func state_handler(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	api_key := params.Get("api_key")
	validate_key(api_key, 0)
}

func handle_state(res http.ResponseWriter, req *http.Request) {
        query, err := url.ParseQuery(req.URL.RawQuery)
        api_key = query["api_key"]
        house_id = query["id"]
	if validate_key(api_key, house_id ) {
	// Convert our array of booleans into a binary representation for http output
	state_value := int64(0)
	for index, value := range Common.state {
		if value == true {
			state_value += int64(math.Exp2(float64(index)))
		}
	}
	fmt.Fprintf(res, "%c", state_value)
	log.Printf("200: Current State: %8b", state_value)
        } 
       else {
http.Error(w, "403 Forbidden : you can't access this resource.", 403)
 }
}

func target_mac_handler(res http.ResponseWriter, req *http.Request) {
	msg := "Welcome to target_mac."
	fmt.Fprintln(res, msg)
	log.Println("200: " + msg)
}

func turn_on(res http.ResponseWriter, req *http.Request) {
	msg := "Welcome to turn_on."
	fmt.Fprintln(res, msg)
	log.Println("200: " + msg)
}

func turn_off(res http.ResponseWriter, req *http.Request) {
	msg := "Welcome to turn_off."
	fmt.Fprintln(res, msg)
	log.Println("200: " + msg)
}

func validate_key(api_key string, house_id int) bool {
        return Common.api_key[house_id] == api_key
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
