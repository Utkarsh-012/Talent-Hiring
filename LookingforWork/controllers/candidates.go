package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"project/LookingforWork/database"
	"project/LookingforWork/modal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidate modal.Candidate
	err := json.NewDecoder(r.Body).Decode(&candidate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform database insertion
	//insertQuery := "INSERT INTO candidate(fullname, about, jobtitle, experience,country) VALUES (?, ?, ?, ?, ?)"
	//_, err = database.Db.Exec(insertQuery, candidate.FullName, candidate.About, candidate.JobTitle, candidate.Experience, candidate.Country)
	insertQuery := "INSERT INTO candidate(fullname, about, jobtitle, experience, country) VALUES (?, ?, ?, ?, ?)"
	_, err = database.Db.Exec(insertQuery, candidate.FullName, candidate.About, candidate.JobTitle, candidate.Experience, candidate.Country)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	candidateID := params["id"]

	// Perform database query
	selectQuery := "SELECT fullname, id, about, jobtitle, experience, country FROM candidate WHERE id = ?"
	row := database.Db.QueryRow(selectQuery, candidateID)

	var candidate modal.Candidate
	err := row.Scan(&candidate.FullName, &candidate.ID, &candidate.About, &candidate.JobTitle, &candidate.Experience, &candidate.Country)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(candidate)
}

func UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	candidateID := params["id"]

	var candidate modal.Candidate
	err := json.NewDecoder(r.Body).Decode(&candidate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform database update
	updateQuery := "UPDATE candidate SET fullname=?, about=?, jobtitle=?, experience=?, country=? WHERE id=?"
	_, err = database.Db.Exec(updateQuery, candidate.FullName, candidate.About, candidate.JobTitle, candidate.Experience, candidate.Country, candidateID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	candidateID := params["id"]

	// Perform database deletion
	deleteQuery := "DELETE FROM candidate WHERE id = ?"
	_, err := database.Db.Exec(deleteQuery, candidateID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
