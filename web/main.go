package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Pong struct {
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func pingHandler(response http.ResponseWriter, request *http.Request) {
	// Set Response Header
	response.Header().Set("Content-Type", "application/json")
	var pong Pong
	pong.Message = "Pong"
	pong.Timestamp = time.Now().UTC().Format(time.RFC850)
	log.Debug(fmt.Sprintf("Index Route Hit from %v", request.RemoteAddr))
	json.NewEncoder(response).Encode(pong)
}

func main()  {
	// Configure Logging
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	// CMD Flags
	portPtr := flag.Int("port", 3000, "Port for server to run on default 3000")
	srvAddressPtr := flag.String("server", "localhost", "Server address default localhost")

	flag.Parse()

	// Init Router
	router := mux.NewRouter()

	// Route handlers
	router.HandleFunc("/ping", pingHandler).Methods("GET")

	// Run Server
	address := fmt.Sprintf("%s:%d", *srvAddressPtr, *portPtr)
	log.Info(fmt.Sprintf("Server starting on http://%s\n", address))
	log.Fatal(http.ListenAndServe(address, router))
}