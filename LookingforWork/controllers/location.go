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

func GetLocations(w http.ResponseWriter, r *http.Request) {
	var locations []modal.Location

	rows, err := database.Db.Query("SELECT * FROM location")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var location modal.Location
		if err := rows.Scan(&location.ID, &location.LocationName, &location.LocationType); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		locations = append(locations, location)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

// GetLocation retrieves a specific location from the database based on the provided ID.
func GetLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	locationID := params["id"]

	var location modal.Location

	row := database.Db.QueryRow("SELECT * FROM location WHERE id = ?", locationID)
	if err := row.Scan(&location.ID, &location.LocationName, &location.LocationType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

// CreateLocation adds a new location to the database.
func CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location modal.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("INSERT INTO location (candidate_id , locationname, locationtype) VALUES (?, ? , ?)", location.CandidateID, location.LocationName, location.LocationType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

// UpdateLocation updates an existing location in the database.
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	locationID := params["id"]

	var location modal.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("UPDATE location SET locationname=?, locationtype=? WHERE id=?", location.LocationName, location.LocationType, locationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Location with ID %s updated successfully", locationID)
}

// DeleteLocation removes a location from the database.
func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	locationID := params["id"]

	_, err := database.Db.Exec("DELETE FROM location WHERE id = ?", locationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Location with ID %s deleted successfully", locationID)
}
