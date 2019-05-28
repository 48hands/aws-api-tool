package services

import (
	"aws-api-tool/models"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type STSService struct{}

type awsSession struct {
	SessionId    string `json:"sessionId"`
	SessionKey   string `json:"sessionKey"`
	SessionToken string `json:"sessionToken"`
}

type signinToken struct {
	value string `json:"SigninToken"`
}

// マネジメントコンソールにログインするための一時的なURLを生成する
func (s STSService) CreateTemporaryConsoleURL(stsClient *sts.STS, federationReq models.FederationRequest) (string, error) {
	issUserURL := "https://www.google.com"
	consoleURL := "https://ap-northeast-1.console.aws.amazon.com/console/home?region=ap-northeast-1"
	signIUnURL := "https://signin.aws.amazon.com/federation"

	signinToken, err := getSignInToken(stsClient, federationReq)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return signIUnURL + "?Action=login" + "&SigninToken=" + url.QueryEscape(signinToken.value) + "&Issuser=" + url.QueryEscape(issUserURL) + "&Destination=" + url.QueryEscape(consoleURL), nil
}

// マネジメントコンソールにサインインするためのトークンを取得する
func getSignInToken(stsClient *sts.STS, federationReq models.FederationRequest) (*signinToken, error) {

	input := sts.AssumeRoleInput{
		DurationSeconds: aws.Int64(int64(federationReq.Durations)),
		RoleArn:         aws.String(federationReq.ARN),
		RoleSessionName: aws.String(federationReq.Username),
	}

	output, err := stsClient.AssumeRole(&input)
	if err != nil {
		log.Fatalln(err)
	}

	jsonByte, _ := json.Marshal(
		awsSession{
			SessionId:    *output.Credentials.AccessKeyId,
			SessionKey:   *output.Credentials.SecretAccessKey,
			SessionToken: *output.Credentials.SessionToken,
		})

	siginInTokenUrl := fmt.Sprintf("https://signin.aws.amazon.com/federation?Action=getSigninToken&SessionType=json&Session=%s", url.QueryEscape(string(jsonByte)))
	res, err := http.Get(siginInTokenUrl)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var token signinToken
	err = json.Unmarshal(body, &token)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
