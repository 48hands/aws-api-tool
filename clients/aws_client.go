package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type AWSApi struct{}

func (awsApi AWSApi) CreateS3Client() *s3.S3 {
	session, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")})

	if err != nil {
		log.Fatal("Error creating session", err)
	}

	return s3.New(session)
}
