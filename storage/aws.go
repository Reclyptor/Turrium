package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
	"time"
	"turrium/env"
)

func Sign(key string, time time.Duration) string {
	if env.AWS_ACCESS_KEY_ID == "" || env.AWS_ACCESS_KEY_SECRET == "" || env.AWS_REGION == "" {
		return ""
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(env.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(env.AWS_ACCESS_KEY_ID, env.AWS_ACCESS_KEY_SECRET, ""),
	})
	if err != nil {
		return ""
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(strings.Split(key, "/")[0]),
		Key:    aws.String(strings.TrimPrefix(key, strings.Split(key, "/")[0])),
	})

	signed, err := req.Presign(time)
	if err != nil {
		return ""
	}

	return signed
}