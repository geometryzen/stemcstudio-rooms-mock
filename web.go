package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()

	router := mux.NewRouter()

	roomsService := NewRoomsService()

	router.HandleFunc("/rooms", createRoom(roomsService)).Methods("POST")
	router.HandleFunc("/rooms/{roomId}", getRoom(roomsService)).Methods("GET")
	router.HandleFunc("/rooms/{roomId}", deleteRoom(roomsService)).Methods("DELETE")

	server := &http.Server{
		Addr:    "0.0.0.0:8082",
		Handler: router,
	}
	fmt.Printf("HTTP server listening at address %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
