package repositories

import (
	"aws-api-tool/models"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type STSRepository struct{}

const (
	issUserUrl = "https://www.google.com"
	consoleUrl = "https://ap-northeast-1.console.aws.amazon.com/console/home?region=ap-northeast-1"
	signinUrl  = "https://signin.aws.amazon.com/federation"

	policy = `
		{
		  "Version": "2012-10-17",
		  "Statement": [
			{
				"Sid": "Stmt1",
				"Effect": "Allow",
				"Action": "s3:*",
				"Resource": "*"
			}
		  ]
		}`
)

func (stsRepository STSRepository) GetFederationToken(stsClient *sts.STS, federatedUserName string) (temporaryCredentials models.TemporaryCredentials, err error) {
	input := &sts.GetFederationTokenInput{
		DurationSeconds: aws.Int64(3600),
		Name:            aws.String(federatedUserName),
		Policy:          aws.String(strings.TrimSpace(policy)),
	}

	output, err := stsClient.GetFederationToken(input)
	if err != nil {
		return models.TemporaryCredentials{}, err
	}

	res, err := accessToSignTokenUrl(convert(output))
	if err != nil {
		return models.TemporaryCredentials{}, err
	}
	defer res.Body.Close()

	signinToken, err := getSigninToken(res)
	if err != nil {
		return models.TemporaryCredentials{}, err
	}

	temporaryCredentials.AccessKey = *output.Credentials.AccessKeyId
	temporaryCredentials.SecretKey = *output.Credentials.SecretAccessKey
	temporaryCredentials.LoginUrl = signinUrl + "?Action=login" + "&SigninToken=" + url.QueryEscape(signinToken) + "&Issuser=" + url.QueryEscape(issUserUrl) + "&Destination=" + url.QueryEscape(consoleUrl)
	return
}

type TemporarySigninToken struct {
	SigninToken string `json:"SigninToken"`
}

func convert(output *sts.GetFederationTokenOutput) models.STSSession {
	var stsSession models.STSSession
	stsSession.SessionId = *output.Credentials.AccessKeyId
	stsSession.SessionKey = *output.Credentials.SecretAccessKey
	stsSession.SessionToken = *output.Credentials.SessionToken

	return stsSession
}

func accessToSignTokenUrl(stsSession models.STSSession) (*http.Response, error) {
	jsonByte, _ := json.Marshal(stsSession)
	jsonStr := string(jsonByte)
	signinTokenUrl := signinUrl + "?Action=getSigninToken" + "&SessionType=json&Session=" + url.QueryEscape(jsonStr)

	res, err := http.Get(signinTokenUrl)
	return res, err
}

func getSigninToken(r *http.Response) (string, error) {
	var temporarySigninToken TemporarySigninToken

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &temporarySigninToken)
	if err != nil {
		return "", err
	}

	return temporarySigninToken.SigninToken, nil
}
