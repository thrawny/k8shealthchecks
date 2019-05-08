package main

import (
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func liveness(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("notlive"); err == nil {
		http.Error(w, "NOTLIVE", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("LIVE"))
}

func readiness(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("notready"); err == nil {
		http.Error(w, "NOTREADY", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("READY"))
}

func index(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	w.Write([]byte(name))
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/liveness", liveness)
	http.HandleFunc("/readiness", readiness)

	log.Info("Starting up on port 8080")

	err := http.ListenAndServe(":8080", handlers.CombinedLoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

}
