package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/drone/location", handleDroneLocation)

	log.Println("Listening on 8080..")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

type DroneLocationRequest struct {
	SectorID float64 `json:"sectorId,string"`
	X        float64 `json:"x,string"`
	Y        float64 `json:"y,string"`
	Z        float64 `json:"z,string"`
	Vel      float64 `json:"vel,string"`
}

type DroneLocationResponse struct {
	Loc float64 `json:"loc,string"`
}

func handleDroneLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var req DroneLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.SectorID <= 0 {
		http.Error(w, "sectorId must be greater than zero", http.StatusBadRequest)
		return
	}

	loc, err := strconv.ParseFloat(fmt.Sprintf("%.2f", req.SectorID*req.X+req.SectorID*req.Y+req.SectorID*req.Z+req.SectorID*req.Vel), 64)
	if err != nil {
		log.Println("ERROR:", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&DroneLocationResponse{Loc: loc}); err != nil {
		log.Println("ERROR:", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
