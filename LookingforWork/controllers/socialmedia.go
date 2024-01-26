package controllers

import (
	"encoding/json"
	"net/http"
	"project/LookingforWork/database"
	"project/LookingforWork/modal"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func GetSocialMedia(w http.ResponseWriter, r *http.Request) {
	var socialMediaEntries []modal.SocialMedia

	rows, err := database.Db.Query("SELECT * FROM socialmedia")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var socialMediaEntry modal.SocialMedia
		if err := rows.Scan(&socialMediaEntry.ID, &socialMediaEntry.Platform, &socialMediaEntry.URL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		socialMediaEntries = append(socialMediaEntries, socialMediaEntry)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(socialMediaEntries)
}

// GetSocialMediaByID retrieves a specific social media entry from the database based on the provided ID.
func GetSocialMediaByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	socialMediaID := params["id"]

	var socialMediaEntry modal.SocialMedia

	row := database.Db.QueryRow("SELECT * FROM socialmedia WHERE id = ?", socialMediaID)
	if err := row.Scan(&socialMediaEntry.ID, &socialMediaEntry.Platform, &socialMediaEntry.URL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(socialMediaEntry)
}

// CreateSocialMedia adds a new social media entry to the database.
func CreateSocialMedia(w http.ResponseWriter, r *http.Request) {
	var socialMediaEntry modal.SocialMedia
	err := json.NewDecoder(r.Body).Decode(&socialMediaEntry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("INSERT INTO socialmedia (candidate_id , platform, url) VALUES (?, ?, ?)", socialMediaEntry.CandidateID, socialMediaEntry.Platform, socialMediaEntry.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(socialMediaEntry)
}

// UpdateSocialMedia updates an existing social media entry in the database.
func UpdateSocialMedia(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	socialMediaID := params["id"]

	var socialMediaEntry modal.SocialMedia
	err := json.NewDecoder(r.Body).Decode(&socialMediaEntry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("UPDATE socialmedia SET platform=?, url=? WHERE id=?", socialMediaEntry.Platform, socialMediaEntry.URL, socialMediaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Social Media Entry with ID %s updated successfully", socialMediaID)
}

// DeleteSocialMedia removes a social media entry from the database.
func DeleteSocialMedia(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	socialMediaID := params["id"]

	_, err := database.Db.Exec("DELETE FROM socialmedia WHERE id = ?", socialMediaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Social Media Entry with ID %s deleted successfully", socialMediaID)
}
