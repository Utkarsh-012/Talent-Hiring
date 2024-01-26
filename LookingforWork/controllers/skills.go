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

func GetSkills(w http.ResponseWriter, r *http.Request) {
	var skills []modal.Skills

	rows, err := database.Db.Query("SELECT * FROM skills")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var skill modal.Skills
		if err := rows.Scan(&skill.ID, &skill.SkillsName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		skills = append(skills, skill)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skills)
}

// GetSkill retrieves a specific skill from the database based on the provided ID.
func GetSkill(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	skillID := params["id"]

	var skill modal.Skills

	row := database.Db.QueryRow("SELECT * FROM skills WHERE id = ?", skillID)
	if err := row.Scan(&skill.ID, &skill.SkillsName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skill)
}

// CreateSkill adds a new skill to the database.
func CreateSkill(w http.ResponseWriter, r *http.Request) {
	var skill modal.Skills
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("INSERT INTO skills (skillsname , candidate_id) VALUES (? , ?)", skill.SkillsName, skill.CandidateID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skill)
}

// UpdateSkill updates an existing skill in the database.
func UpdateSkill(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	skillID := params["id"]

	var skill modal.Skills
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.Db.Exec("UPDATE skills SET skillsname=? WHERE id=?", skill.SkillsName, skillID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Skill with ID %s updated successfully", skillID)
}

// DeleteSkill removes a skill from the database.
func DeleteSkill(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	skillID := params["id"]

	_, err := database.Db.Exec("DELETE FROM skills WHERE id = ?", skillID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Skill with ID %s deleted successfully", skillID)
}
