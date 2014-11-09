// Currently everything for ether_housed is in the main package
package main

import (
	"fmt"
	"log"
	"log/syslog"
	"math"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"github.com/bradfitz/gomemcache/memcache"
)

const NUM_HOUSES = 8

// common is a struct to store the global state and config
type common struct {
	lock       sync.RWMutex
	state      []bool
	api_key    []string
	target_mac []string
}

var Common = new(common)

// Get is a method on the common struct to safely read
// the state id of a particular house
func (c *common) Get(id int) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	d := c.state[id]
	return d
}

// Set will lock and set the state of a house
func (c *common) Set(id int, d bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.state[id] = d
}

// load_existing_state pulls state from an external datastore
func load_existing_state() {
	// BUG: Get state out of memcache
	Common.state = []bool{false, false, false, false, false, false, false, false}
	return
}

// load_target_macs gets the config for which MAC addresses are associated with each house
// from the environment
func load_target_macs() {
	Common.target_mac = []string{"", "", "", "", "", "", "", ""}
	for i := 0; i < NUM_HOUSES; i++ {
		Common.target_mac[i] = strings.TrimSpace(os.Getenv("MAC" + strconv.Itoa(i)))
		if Common.target_mac[i] == "" {
			log.Println("WARNING: Didn't get an MAC for " + strconv.Itoa(i) + ".")
		}
		log.Println("INFO: MAC for " + strconv.Itoa(i) + " is " + Common.target_mac[i])
	}
}

// load_api_keys reads in an APIKEYN environment variable for each houseid
func load_api_keys() {
	Common.api_key = []string{"", "", "", "", "", "", "", ""}
	for i := 0; i < NUM_HOUSES; i++ {
		Common.api_key[i] = strings.TrimSpace(os.Getenv("APIKEY" + strconv.Itoa(i)))
		if Common.api_key[i] == "" {
			log.Println("WARNING: Didn't get an API key for " + strconv.Itoa(i) + ".")
		}
		log.Println("INFO: API Key for " + strconv.Itoa(i) + " is " + Common.api_key[i])
	}
}

func setup_logging() {
	logwriter, e := syslog.New(syslog.LOG_NOTICE, "ether_housed")
	if e == nil {
		log.SetOutput(logwriter)
	}
}

var chttp = http.NewServeMux()

func initialize_memcached() {
	servers := os.Getenv("MEMCACHEDCLOUD_SERVERS")
	username := os.Getenv("MEMCACHEDCLOUD_USERNAME")
	password := os.Getenv("MEMCACHEDCLOUD_PASSWORD")
	if servers != "" && username != "" && password != "" {
		log.Println("Read memcache config from env")
	} else {
		log.Println("Failed to read MEMCACHEDCLOUD Variables. Trying localhost next")
		mc := memcache.New("127.0.0.1:11211")
		fmt.Println(mc)
		log.Println("Failed to read memcache from env. Going without it.")
	}
}

func main() {
	initialize_memcached()

	load_existing_state()
	load_api_keys()
	load_target_macs()
	http.HandleFunc("/", usage)
	http.HandleFunc("/on", turn_on)
	http.HandleFunc("/off", turn_off)
	http.HandleFunc("/state", handle_state)
	http.HandleFunc("/info", handle_info)
	http.HandleFunc("/target_mac", target_mac_handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Print("No PORT variable. Defaulting to 3000")
	}
	log.Print("listening on " + port + "...")
	chttp.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func usage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		chttp.ServeHTTP(res, req)
	} else {
		msg := "Welcome to ether_house.\n"
		msg += "Source code: https://github.com/solarkennedy/ether_housed \n"
		msg += "Client code: https://github.com/solarkennedy/ether_house \n"
		fmt.Fprintln(res, msg)
		log.Println("200: " + req.URL.Path)
	}
}

// boolarraytoint converts our array of booleans into a binary representation for http output
func boolarraytoint(bool_array []bool) (out int) {
	for index, value := range bool_array {
		if value == true {
			out += int(int64(math.Exp2(float64(index))))
		}
	}
	return out
}

func mactobinary(mac string) (output []byte) {
	output, err := net.ParseMAC(mac)
	if err != nil {
		log.Printf("Error parsing mac: %v, output: %v, error: %v", mac, output, err)
	}
	return output
}

func get_state_as_int() (state_int int) {
	Common.lock.Lock()
	defer Common.lock.Unlock()
	return boolarraytoint(Common.state)
}

func handle_state(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	api_key := query.Get("api_key")
	house_id_string := query.Get("id")
	house_id, _ := strconv.ParseInt(house_id_string, 0, 64)
	if validate_key(api_key, int(house_id)) {
		state_value := get_state_as_int()
		tmp_bytes := []byte{byte(state_value)}
		res.Write(tmp_bytes)
		log.Printf("200: Current State: %08b", state_value)
	} else {
		http.Error(res, "403 Forbidden : you can't access this resource.", 403)
		log.Printf("403: /state from %v, using api key %v", house_id, api_key)
	}
}

func handle_info(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	api_key := query.Get("api_key")
	house_id_string := query.Get("id")
	house_id, _ := strconv.ParseInt(house_id_string, 0, 64)
	if validate_key(api_key, int(house_id)) {
		state_value := get_state_as_int()
		target_mac := Common.target_mac[house_id]
		fmt.Fprintf(res, "Hi!!! Curious about how this works? Here is some debug info.\n\n\n")
		fmt.Fprintf(res, "Information on house_id: %v\n", house_id)
		fmt.Fprintf(res, "Current state: "+"%08b (%v)\n", state_value, state_value)
		fmt.Fprintf(res, "Target MAC Address: %v\n\n\n", target_mac)
		fmt.Fprintf(res, "Server Source code: https://github.com/solarkennedy/ether_housed \n")
		fmt.Fprintf(res, "Client code: https://github.com/solarkennedy/ether_house \n")
		log.Printf("200: /info for %v", house_id)
	} else {
		http.Error(res, "403 Forbidden : you can't access this resource.", 403)
		log.Printf("403: /info from %v, using api key %v", house_id, api_key)
	}
}

func target_mac_handler(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	api_key := query.Get("api_key")
	house_id_string := query.Get("id")
	house_id, _ := strconv.ParseInt(house_id_string, 0, 64)
	if validate_key(api_key, int(house_id)) {
		target_mac := Common.target_mac[house_id]
		target_mac_binary := mactobinary(target_mac)
		target_mac_string := string(target_mac_binary[:6])
		fmt.Fprintf(res, target_mac_string)
		log.Printf("200: target_mac: %v ", target_mac)
	} else {
		http.Error(res, "403 Forbidden : you can't access this resource.", 403)
		log.Printf("403: /state from %v, using api key %v", house_id, api_key)
	}
}

func turn_on(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	api_key := query.Get("api_key")
	house_id_string := query.Get("id")
	house_id, _ := strconv.ParseInt(house_id_string, 0, 64)
	if validate_key(api_key, int(house_id)) {
		Common.Set(int(house_id), true)
		log.Printf("200: turn_on: %v", house_id)
	} else {
		http.Error(res, "403 Forbidden : you can't access this resource.", 403)
		log.Printf("403: /on from %v, using api key %v", house_id, api_key)
	}
}

func turn_off(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	api_key := query.Get("api_key")
	house_id_string := query.Get("id")
	house_id, _ := strconv.ParseInt(house_id_string, 0, 64)
	if validate_key(api_key, int(house_id)) {
		Common.Set(int(house_id), false)
		log.Printf("200: turn_off: %v", house_id)
	} else {
		http.Error(res, "403 Forbidden : you can't access this resource.", 403)
		log.Printf("403: /off from %v, using api key %v", house_id, api_key)
	}
}

// validate_key ensures that the provided key matches the one stored for that house_id
func validate_key(api_key string, house_id int) bool {
	return Common.api_key[house_id] == api_key
}
