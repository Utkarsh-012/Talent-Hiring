package controllers

import (
	"encoding/json"

	"fmt"
	"net/http"

	"project/LookingtoHire/modal"
)

func getJobFromLookingForWork() ([]modal.Job, error) {
	response, err := http.Get("http://localhost:9090/job")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch candidates. Status: %d", response.StatusCode)
	}
	var jobs []modal.Job
	if err := json.NewDecoder(response.Body).Decode(&jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	jobs, err := getJobFromLookingForWork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process candidates and send the response
	json.NewEncoder(w).Encode(jobs)
}
