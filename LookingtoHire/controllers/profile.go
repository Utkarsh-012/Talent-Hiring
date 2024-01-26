package controllers

import (
	"encoding/json"

	"fmt"
	"net/http"

	"project/LookingtoHire/modal"
)

func getProfileFromLookingForWork() ([]modal.Profile, error) {
	response, err := http.Get("http://localhost:9090/profile")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch candidates. Status: %d", response.StatusCode)
	}
	var profiles []modal.Profile

	if err := json.NewDecoder(response.Body).Decode(&profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	profiles, err := getProfileFromLookingForWork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process candidates and send the response
	json.NewEncoder(w).Encode(profiles)
}
