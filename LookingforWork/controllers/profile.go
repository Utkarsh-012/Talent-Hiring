package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/LookingforWork/database"
	"project/LookingforWork/modal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	var profiles []modal.Profile

	rows, err := database.Db.Query("SELECT * FROM profiles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var profile modal.Profile
		if err := rows.Scan(&profile.ID, &profile.ProfileLink); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		profiles = append(profiles, profile)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

// GetProfile retrieves a specific profile from the database based on the provided ID.
func GetProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileID := params["id"]

	var profile modal.Profile

	row := database.Db.QueryRow("SELECT * FROM profiles WHERE id = ?", profileID)
	if err := row.Scan(&profile.ID, &profile.ProfileLink); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// CreateProfile adds a new profile to the database.
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profile modal.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("INSERT INTO profile (candidate_id , profilelink) VALUES (?,?)", profile.CandidateID, profile.ProfileLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// UpdateProfile updates an existing profile in the database.
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileID := params["id"]

	var profile modal.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("UPDATE profiles SET profilelink=? WHERE id=?", profile.ProfileLink, profileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Profile with ID %s updated successfully", profileID)
}

// DeleteProfile removes a profile from the database.
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileID := params["id"]

	_, err := database.Db.Exec("DELETE FROM profiles WHERE id = ?", profileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Profile with ID %s deleted successfully", profileID)
}
