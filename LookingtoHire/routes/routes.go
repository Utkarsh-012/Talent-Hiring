package routes

import (
	"project/LookingtoHire/controllers"

	"github.com/gorilla/mux"
)

var RegisterTalentroutes = func(router *mux.Router) {

	router.HandleFunc("/candidate", controllers.GetCandidate).Methods("GET")

	router.HandleFunc("/company", controllers.GetCompany).Methods("GET")
	router.HandleFunc("/company", controllers.AddCompany).Methods("POST")
	router.HandleFunc("/company", controllers.UpdateCompany).Methods("PUT")
	router.HandleFunc("/company", controllers.DeleteCompany).Methods("DELETE")

	router.HandleFunc("/profile", controllers.GetProfile).Methods("GET")

	router.HandleFunc("/location", controllers.GetLocation).Methods("GET")

	router.HandleFunc("/job", controllers.GetJob).Methods("GET")

	router.HandleFunc("/skills", controllers.GetSkills).Methods("GET")

	router.HandleFunc("/socialmedia", controllers.GetSocialMedia).Methods("GET")

}
