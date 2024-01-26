package controllers

import (
	"encoding/json"

	"fmt"
	"net/http"

	"project/LookingtoHire/modal"
)

func getLocationFromLookingForWork() ([]modal.Location, error) {
	response, err := http.Get("http://localhost:9090/location")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch candidates. Status: %d", response.StatusCode)
	}
	var locations []modal.Location

	if err := json.NewDecoder(response.Body).Decode(&locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	locations, err := getLocationFromLookingForWork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process candidates and send the response
	json.NewEncoder(w).Encode(locations)
}
