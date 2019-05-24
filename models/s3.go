package models

import "time"

type S3Bucket struct {
	Name string `json:"name"`
	CreationTime time.Time `json:"creation_time"`
}

type S3BucketDetail struct {
	Name string `json:"name"`
	Objects []S3Object `json:"objects"`
}

type S3Object struct {
	Key string `json:"key"`
	LastModified time.Time `json:"last_modified"`
	Size int64 `json:"size"`
}