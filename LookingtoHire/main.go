package main

import (
	"log"
	"net/http"
	"project/LookingtoHire/database"
	"project/LookingtoHire/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTalentroutes(r)
	database.InitDB()
	// database.CreateTableCompany()

	log.Fatal(http.ListenAndServe("localhost:9091", r))
}
