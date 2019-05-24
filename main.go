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
	s3Client := awsApi.CreateS3Client()
	s3Controller := controllers.S3Controller{}

	router.HandleFunc("/s3buckets", s3Controller.GetBuckets(s3Client)).Methods("GET")
	router.HandleFunc("/s3buckets/{bucketName}", s3Controller.GetBucket(s3Client)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
