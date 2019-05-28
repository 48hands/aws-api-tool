package main

import (
	"aws-api-tool/clients"
	"aws-api-tool/controllers"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"

	"log"
	"net/http"
)

func init() {
	gotenv.Load()
}

func main() {
	router := mux.NewRouter()

	awsApi := clients.AWSApi{}
	stsClient := awsApi.CreateSTSClient()
	stsController := controllers.STSController{}

	router.HandleFunc("/federation", stsController.CreateTemporaryConsoleURL(stsClient)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
