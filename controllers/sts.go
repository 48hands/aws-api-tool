package controllers

import (
	"aws-api-tool/models"
	"aws-api-tool/repositories"
	"aws-api-tool/utils"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/gorilla/mux"
	"net/http"
)

type STSController struct{}

func (c STSController) CreateTemporaryURL(stsClient *sts.STS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		federatedUserName := mux.Vars(r)["username"]
		stsRepo := repositories.STSRepository{}
		temporaryCredentials, err := stsRepo.GetFederationToken(stsClient, federatedUserName)
		if err != nil {
			var error models.Error
			error.Message = "Error Occurred!"
			utils.SendError(w, http.StatusServiceUnavailable, error)
		}
		utils.SendSuccess(w, temporaryCredentials)
	}
}
