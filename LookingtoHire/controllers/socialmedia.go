package controllers

import (
	"encoding/json"

	"fmt"
	"net/http"

	"project/LookingtoHire/modal"
)

func getSocialMediaFromLookingForWork() ([]modal.SocialMedia, error) {
	response, err := http.Get("http://localhost:9090/socialmedia")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch candidates. Status: %d", response.StatusCode)
	}
	var socialmedia []modal.SocialMedia

	if err := json.NewDecoder(response.Body).Decode(&socialmedia); err != nil {
		return nil, err
	}

	return socialmedia, nil
}

func GetSocialMedia(w http.ResponseWriter, r *http.Request) {
	socialmedia, err := getSocialMediaFromLookingForWork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process candidates and send the response
	json.NewEncoder(w).Encode(socialmedia)
}
