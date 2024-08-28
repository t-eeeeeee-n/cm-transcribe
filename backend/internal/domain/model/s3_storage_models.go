package model

import "fmt"

type S3File struct {
	FilePath   string
	BucketName string
	KeyPrefix  string
}

func NewS3File(path, bucketName, keyPrefix string) *S3File {
	return &S3File{
		FilePath:   path,
		BucketName: bucketName,
		KeyPrefix:  keyPrefix,
	}
}

func (s *S3File) Validate() error {
	if s.FilePath == "" || s.BucketName == "" {
		return fmt.Errorf("FilePath and BucketName are required")
	}
	return nil
}
