package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
)

var (
	Debug      bool
	S3StateKey string

	s3StateBucket string
)

var (
	Sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	EcsClient            = ecs.New(Sess)
	ResourceGroupsClient = resourcegroupstaggingapi.New(Sess)
	S3Client             = s3.New(Sess)
	STSClient            = sts.New(Sess)
	RDSClient            = rds.New(Sess)
)

func init() {
	Debug = getenv("DEBUG", true).(bool)
	S3StateKey = getenv("AWS_SWITCH_STATE_KEY", "aws-switch").(string)

	s3StateBucket = getenv("AWS_SWITCH_STATE_BUCKET", "").(string)
}

func GetS3StateBucket() string {
	if s3StateBucket == "" {
		output, err := STSClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
		if err != nil {
			fmt.Printf("Error getting s3 state bucket: %s", err)
			os.Exit(1)
		}
		s3StateBucket = fmt.Sprintf("%s-aws-switch", *output.Account)
	}
	return s3StateBucket
}

func getenv(key string, fallback interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	switch fallback.(type) {
	case string:
		v := os.Getenv(key)
		if len(value) == 0 {
			return fallback
		}
		return v
	case int:
		s := os.Getenv(key)
		v, err := strconv.Atoi(s)
		if err != nil {
			return fallback
		}
		return v

	case bool:
		s := os.Getenv(key)
		v, err := strconv.ParseBool(s)
		if err != nil {
			return fallback
		}
		return v
	default:
		return value
	}
}
