package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getTime).Queries("tz", "{tz}")
	router.HandleFunc("/api/time", getTime)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
