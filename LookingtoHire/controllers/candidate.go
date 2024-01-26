package controllers

import (
	"encoding/json"
	"project/LookingforWork/modal"

	"fmt"
	"net/http"
	// "project/LookingtoHire/modal"
)

func getCandidatesFromLookingForWork() ([]modal.Candidate, error) {
	response, err := http.Get("http://localhost:9090/candidates")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch candidates. Status: %d", response.StatusCode)
	}
	var candidates []modal.Candidate

	if err := json.NewDecoder(response.Body).Decode(&candidates); err != nil {
		return nil, err
	}

	return candidates, nil
}

func GetCandidate(w http.ResponseWriter, r *http.Request) {
	candidates, err := getCandidatesFromLookingForWork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process candidates and send the response
	json.NewEncoder(w).Encode(candidates)
}
