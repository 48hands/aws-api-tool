package models

type FederationRequest struct {
	Username  string `json:"username"`
	ARN       string `json:"arn"`
	Durations int    `json:"durations"`
}

type FederationResponse struct{}
