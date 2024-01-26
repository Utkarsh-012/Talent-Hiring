package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"project/LookingtoHire/modal"
	"strconv"

	"github.com/gorilla/mux"
)

var NewCompany []*modal.Company

func GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(NewCompany)

}

func AddCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Company modal.Company
	_ = json.NewDecoder(r.Body).Decode(&Company)
	Company.ID = strconv.Itoa(rand.Intn(1000000))
	NewCompany = append(NewCompany, &Company)
	json.NewEncoder(w).Encode(Company)

}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range NewCompany {

		if item.ID == params["id"] {
			NewCompany = append(NewCompany[:index], NewCompany[index+1])
			break
		}

	}
	json.NewEncoder(w).Encode(NewCompany)

}
func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range NewCompany {
		if item.ID == params["id"] {
			NewCompany = append(NewCompany[:index], NewCompany[index+1])
			var Company modal.Company
			_ = json.NewDecoder(r.Body).Decode(&Company)
			Company.ID = strconv.Itoa(rand.Intn(1000000))
			NewCompany = append(NewCompany, &Company)
			json.NewEncoder(w).Encode(Company)

		}
	}

}
