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

func GetJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []modal.Job

	rows, err := database.Db.Query("SELECT * FROM jobs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var job modal.Job
		if err := rows.Scan(&job.ID, &job.JobName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jobs = append(jobs, job)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

// GetJob retrieves a specific job from the database based on the provided ID.
func GetJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jobID := params["id"]

	var job modal.Job

	row := database.Db.QueryRow("SELECT * FROM job WHERE jobID = ?", jobID)
	if err := row.Scan(&job.ID, &job.JobName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

// CreateJob adds a new job to the database.
func CreateJob(w http.ResponseWriter, r *http.Request) {
	var job modal.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("INSERT INTO job (candidate_id, jobName) VALUES (?,?)", job.CandidateID, job.JobName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

// UpdateJob updates an existing job in the database.
func UpdateJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jobID := params["id"]

	var job modal.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("UPDATE jobs SET jobName=? WHERE jobID=?", job.JobName, jobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Job with ID %s updated successfully", jobID)
}

// DeleteJob removes a job from the database.
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jobID := params["id"]

	_, err := database.Db.Exec("DELETE FROM job WHERE jobID = ?", jobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Job with ID %s deleted successfully", jobID)
}
