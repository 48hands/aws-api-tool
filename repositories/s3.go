package repositories

import (
	"aws-api-tool/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type S3Repository struct{}

func (repo S3Repository) GetBuckets(s3Client *s3.S3) ([]models.S3Bucket, error) {
	input := &s3.ListBucketsInput{}
	result, err := s3Client.ListBuckets(input)
	var s3Buckets []models.S3Bucket

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatalln(aerr.Error())
				return []models.S3Bucket{}, aerr
			}
		}

		return []models.S3Bucket{}, err
	}

	for _, b := range result.Buckets {
		var s3Bucket models.S3Bucket
		s3Bucket.Name = *b.Name
		s3Bucket.CreationTime = *b.CreationDate
		s3Buckets = append(s3Buckets, s3Bucket)
	}

	return s3Buckets, nil

}

func (repo S3Repository) GetBucket(s3Client *s3.S3, bucketName string) (models.S3BucketDetail, error) {
	var s3Objects []models.S3Object
	var s3Bucket models.S3BucketDetail

	input := &s3.ListObjectsInput{Bucket: aws.String(bucketName)}
	objects, err := s3Client.ListObjects(input)
	if err != nil {
		log.Fatal(err)
		return models.S3BucketDetail{}, err
	}

	for _, o := range objects.Contents {
		var s3Object models.S3Object
		s3Object.Key = *o.Key
		s3Object.LastModified = *o.LastModified
		s3Object.Size = *o.Size

		s3Objects = append(s3Objects, s3Object)
	}

	s3Bucket.Name = bucketName
	s3Bucket.Objects = s3Objects

	return s3Bucket, nil
}
