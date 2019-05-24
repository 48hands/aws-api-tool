package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
)

type AWSApi struct{}

func createSession() *session.Session {
	session, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")})

	if err != nil {
		log.Fatal("Error creating session", err)
	}
	return session
}

func (awsApi AWSApi) CreateS3Client() *s3.S3 {
	return s3.New(createSession())
}

func (awsApi AWSApi) CreateSTSClient() *sts.STS {
	return sts.New(createSession())
}
