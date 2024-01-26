package routes

import (
	"project/LookingforWork/controllers"

	"github.com/gorilla/mux"
)

var RegisterTalentRoutes = func(router *mux.Router) {

	router.HandleFunc("/candidate", controllers.CreateCandidate).Methods("POST")
	router.HandleFunc("/candidate/{id}", controllers.GetCandidate).Methods("GET")
	router.HandleFunc("/candidate/{id}", controllers.UpdateCandidate).Methods("PUT")
	router.HandleFunc("/candidate/{id}", controllers.DeleteCandidate).Methods("DELETE")

	router.HandleFunc("/profile", controllers.GetProfiles).Methods("GET")
	router.HandleFunc("/profile/{id}", controllers.GetProfile).Methods("GET")
	router.HandleFunc("/profile", controllers.CreateProfile).Methods("POST")
	router.HandleFunc("/profile/{id}", controllers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/profile/{id}", controllers.DeleteProfile).Methods("DELETE")

	router.HandleFunc("/location", controllers.GetLocations).Methods("GET")
	router.HandleFunc("/location/{id}", controllers.GetLocation).Methods("GET")
	router.HandleFunc("/location", controllers.CreateLocation).Methods("POST")
	router.HandleFunc("/location/{id}", controllers.UpdateLocation).Methods("PUT")
	router.HandleFunc("/location/{id}", controllers.DeleteLocation).Methods("DELETE")

	router.HandleFunc("/skill", controllers.GetSkills).Methods("GET")
	router.HandleFunc("/skill/{id}", controllers.GetSkill).Methods("GET")
	router.HandleFunc("/skill", controllers.CreateSkill).Methods("POST")
	router.HandleFunc("/skill/{id}", controllers.UpdateSkill).Methods("PUT")
	router.HandleFunc("/skill/{id}", controllers.DeleteSkill).Methods("DELETE")

	router.HandleFunc("/socialmedia", controllers.GetSocialMedia).Methods("GET")
	router.HandleFunc("/socialmedia/{id}", controllers.GetSocialMediaByID).Methods("GET")
	router.HandleFunc("/socialmedia", controllers.CreateSocialMedia).Methods("POST")
	router.HandleFunc("/socialmedia/{id}", controllers.UpdateSocialMedia).Methods("PUT")
	router.HandleFunc("/socialmedia/{id}", controllers.DeleteSocialMedia).Methods("DELETE")

	router.HandleFunc("/job", controllers.CreateJob).Methods("POST")
	router.HandleFunc("/job/{id}", controllers.GetJob).Methods("GET")
	router.HandleFunc("/job/{id}", controllers.UpdateJob).Methods("PUT")
	router.HandleFunc("/job/{id}", controllers.DeleteJob).Methods("DELETE")
	router.HandleFunc("/job", controllers.GetJobs).Methods("GET")
}
