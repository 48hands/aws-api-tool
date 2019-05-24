package controllers

import (
	"aws-api-tool/models"
	"aws-api-tool/repositories"
	"aws-api-tool/utils"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"net/http"
)

type S3Controller struct{}

func (c S3Controller) GetBuckets(s3Client *s3.S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s3Repo := repositories.S3Repository{}
		s3Buckets, err := s3Repo.GetBuckets(s3Client)
		if err != nil {
			var error models.Error
			error.Message = "Error Occurred!"
			utils.SendError(w, http.StatusServiceUnavailable, error)
		}
		utils.SendSuccess(w, s3Buckets)
	}
}

func (c S3Controller) GetBucket(s3Client *s3.S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bucketName := mux.Vars(r)["bucketName"]
		s3Repo := repositories.S3Repository{}
		s3Bucket, err := s3Repo.GetBucket(s3Client, bucketName)

		if err != nil {
			var error models.Error
			error.Message = "Error Occurred!"
			utils.SendError(w, http.StatusServiceUnavailable, error)
		}

		utils.SendSuccess(w, s3Bucket)
	}
}
