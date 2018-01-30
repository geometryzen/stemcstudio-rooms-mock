package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type createRoomRequest struct {
	Owner       string  `json:"owner"`
	Description string  `json:"description"`
	Public      bool    `json:"public"`
	Expire      float64 `json:"expire"`
}

type roomResponse struct {
	ID          string `json:"id"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

func createRoom(service RoomsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRoomRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		fmt.Printf("Creating Room for Owner %s\n", req.Owner)
		args := RoomParams{Description: req.Description, Expire: req.Expire, Owner: req.Owner, Public: req.Public}
		room, err := service.CreateRoom(&args)
		if err != nil {
			fmt.Println("err  : ", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		response := roomResponse{ID: room.ID, Owner: room.Owner, Description: room.Description, Public: room.Public}
		json.NewEncoder(w).Encode(&response)
	}
}

func getRoom(service RoomsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomID := vars["roomId"]
		fmt.Printf("GetRoom %s\n", roomID)
		room, err := service.GetRoom(roomID)
		if err != nil {
			fmt.Println("err  : ", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		response := roomResponse{ID: room.ID, Owner: room.Owner, Description: room.Description, Public: room.Public}
		json.NewEncoder(w).Encode(&response)
	}
}

func deleteRoom(service RoomsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomID := vars["roomId"]
		fmt.Printf("DestroyRoom %s\n", roomID)
		_, err := service.DestroyRoom(roomID)
		if err != nil {
			fmt.Println("err  : ", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		http.Error(w, "OK", http.StatusOK)
	}
}
