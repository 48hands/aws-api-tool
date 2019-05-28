package controllers

import (
	"aws-api-tool/models"
	"aws-api-tool/services"
	"aws-api-tool/utils"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/sts"
	"net/http"
)

type STSController struct{}

func (c STSController) CreateTemporaryConsoleURL(stsClient *sts.STS) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var federationReq models.FederationRequest
		var error models.Error
		s := services.STSService{}

		err := json.NewDecoder(r.Body).Decode(&federationReq)
		if err != nil {
			error.Message = "Error Occurred!"
			utils.SendError(w, http.StatusBadRequest, error)
		}

		url, err := s.CreateTemporaryConsoleURL(stsClient, federationReq)
		if err != nil {
			error.Message = "Error Occurred!"
			utils.SendError(w, http.StatusServiceUnavailable, error)
		}

		utils.SendSuccess(w, url)
	}
}
