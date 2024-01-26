// candidate.go
package main

import (
	"log"
	"net/http"
	"project/LookingforWork/database"
	"project/LookingforWork/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTalentRoutes(r)

	database.InitDB()
	database.CreateTable()
	database.CreateTableJob()
	database.CreateTableProfile()
	database.CreateTableSkills()
	database.CreateTableSocialMedia()
	database.CreateTableLocation()
	// Implement routes for UpdateCandidate, DeleteCandidate, GetAllCandidates, etc.

	// Similarly, set up routes for jobs.

	log.Fatal(http.ListenAndServe(":8080", r))
}
