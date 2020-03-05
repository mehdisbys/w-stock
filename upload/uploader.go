package upload

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	AWS_S3_REGION = "eu-west-1"
	AWS_S3_BUCKET = "wzyd"
)

var sess = connectAWS()

func connectAWS() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_S3_REGION),
		},
	)
	if err != nil {
		panic(err)
	}
	return sess
}

func Uploader(filename string, file []byte) (string, error) {
	uploader := s3manager.NewUploader(sess)

	output, err := uploader.Upload(&s3manager.UploadInput{
		ContentType: aws.String("image/jpeg"),
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(AWS_S3_BUCKET),
		Key:         aws.String("imgs/" + filename),
		Body:        bytes.NewReader(file),
	})

	if err != nil {
		return "", err
	}

	return output.Location, err
}
