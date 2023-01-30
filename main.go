package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api", ApiHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static/"))))

	r.PathPrefix("/").HandlerFunc(CatchAllHandler)

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:1205",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	values := map[string]string{"status": "ok", "message": ""}

	jsonValue, _ := json.Marshal(values)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonValue)
}

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/build/index.html")
}
