package controllers

import (
	"encoding/json"

	"fmt"
	"net/http"

	"project/LookingtoHire/modal"
)

func getSkillsFromLookingForWork() ([]modal.Skills, error) {
	response, err := http.Get("http://localhost:9090/skills")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch candidates. Status: %d", response.StatusCode)
	}
	var skills []modal.Skills

	if err := json.NewDecoder(response.Body).Decode(&skills); err != nil {
		return nil, err
	}

	return skills, nil
}

func GetSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := getSkillsFromLookingForWork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process candidates and send the response
	json.NewEncoder(w).Encode(skills)
}
