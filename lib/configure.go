package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/chrispruitt/aws-switch/config"
	"github.com/chrispruitt/aws-switch/state"
)

func Configure() (string, error) {
	bucket := config.GetS3StateBucket()
	_, err := config.S3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	state.InitalizeState()
	return bucket, err
}
