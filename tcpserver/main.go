package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var sport = ""

func loggingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	loggingHandler(w, r)
	io.WriteString(w, "Site up and running")
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	loggingHandler(w, r)
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "Unknown"
	}
	io.WriteString(w, fmt.Sprintf("Golang server Host: %s Port%s", hostname, sport))
}

func handleError(w http.ResponseWriter, r *http.Request) {
	loggingHandler(w, r)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func main() {
	sport = os.Getenv("TCP_PORT")

	port, err := strconv.Atoi(sport)

	if port <= 0 || port > 65535 {
		log.Println("Unable to set desired port number, defaulting to 8080")
		sport = "8080"
	}

	// Set routing rules
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/status", handleStatus)
	http.HandleFunc("/error", handleError)

	//Use the default DefaultServeMux.
	sport = ":" + sport
	log.Println("Server starting and listening on port", sport)
	err = http.ListenAndServe(sport, nil)
	if err != nil {
		log.Fatal(err)
	}
}
