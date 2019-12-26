package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

type Pong struct {
	Message string `json:"message"`
	Timestamp string `json:"now"`
	UpTime string `json:"up_time""`
}

var startedAt time.Time

func pingHandler(response http.ResponseWriter, request *http.Request) {
	var pong Pong
	pong.Message = "Pong"
	pong.Timestamp = time.Now().UTC().Format(time.RFC850)
	pong.UpTime = fmt.Sprintf("%v", time.Now().UTC().Sub(startedAt.UTC()))
	json.NewEncoder(response).Encode(pong)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		host, _, err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {
			host = r.RemoteAddr
		}
		log.WithFields(log.Fields{
			"client_ip": host,
			"uri":   r.RequestURI,
			"user-agent": r.UserAgent(),
			"referer": r.Referer(),
		}).Info(fmt.Sprintf("Request Made for %s", r.RequestURI))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main()  {
	// set start time
	startedAt = time.Now()
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

	// Setup logged Router
	router.Use(loggerMiddleware)
	router.Use(jsonMiddleware)

	// Run Server
	address := fmt.Sprintf("%s:%d", *srvAddressPtr, *portPtr)
	log.Info(fmt.Sprintf("Server starting on http://%s\n", address))
	log.Fatal(http.ListenAndServe(address, handlers.CORS()(router)))
}