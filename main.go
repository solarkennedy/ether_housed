package main

import (
	"fmt"
	"log"
	"log/syslog"
	"math"
	"net"
	"net/http"
	"net/url"
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

func load_existing_state() {
	// TODO: Get state out of memcache
	Common.state = []bool{false, false, false, false, true, true, true, true}
	return
}

func load_target_macs() {
	// We get the config for which MAC addresses are associated with each house
	// From the environment
	Common.target_mac = []string{"", "", "", "", "", "", "", ""}
	for i := 0; i < NUM_HOUSES; i++ {
		Common.target_mac[i] = os.Getenv("MAC" + strconv.Itoa(i))
		if Common.target_mac[i] == "" {
			log.Println("WARNING: Didn't get an MAC for " + strconv.Itoa(i) + ".")
		}
		log.Println("INFO: MAC for " + strconv.Itoa(i) + " is " + Common.target_mac[i])
	}
}

func load_api_keys() {
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
	load_api_keys()
	load_target_macs()
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
	msg := "Welcome to ether_house.\n"
	msg += "Source code: https://github.com/solarkennedy/ether_housed \n"
	msg += "Client code: https://github.com/solarkennedy/ether_house \n"
	fmt.Fprintln(res, msg)
	log.Println("200: " + req.URL.Path)
}

func boolarraytoint(bool_array []bool) (the_int int64) {
	// Convert our array of booleans into a binary representation for http output
	for index, value := range bool_array {
		if value == true {
			the_int += int64(math.Exp2(float64(index)))
		}
	}
	return the_int
}

func mactobinary(mac string) (output []byte) {
	output, _ = net.ParseMAC(mac)
	return output
}

func handle_state(res http.ResponseWriter, req *http.Request) {
	query, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		http.Error(res, "500: Couldn't parse query", 500)
		log.Printf("500: Error on %v", req.URL.RawQuery)
	}
	api_key := query["api_key"][0]
	house_id_string := query["id"][0]
	house_id, _ := strconv.ParseInt(house_id_string, 0, 64)
	if validate_key(api_key, int(house_id)) {
		state_value := boolarraytoint(Common.state)
		fmt.Fprintf(res, "%c", state_value)
		log.Printf("200: Current State: %8b", state_value)
	} else {
		http.Error(res, "403 Forbidden : you can't access this resource.", 403)
		log.Printf("403: /state from %v, using api key %v", house_id, api_key)
	}
}

func target_mac_handler(res http.ResponseWriter, req *http.Request) {
	query, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		http.Error(res, "500: Couldn't parse query", 500)
		log.Printf("500: Error on %v", req.URL.RawQuery)
	}
	api_key := query["api_key"][0]
	house_id_string := query["id"][0]
	house_id, _ := strconv.ParseInt(house_id_string, 0, 64)
	if validate_key(api_key, int(house_id)) {
		target_mac := Common.target_mac[house_id]
		target_mac_binary := mactobinary(target_mac)
		target_mac_string := string(target_mac_binary[:6])
		fmt.Fprintf(res, target_mac_string)
		log.Printf("200: target_mac: ", target_mac)
	} else {
		http.Error(res, "403 Forbidden : you can't access this resource.", 403)
		log.Printf("403: /state from %v, using api key %v", house_id, api_key)
	}
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
