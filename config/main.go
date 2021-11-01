package config

import (
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	Debug         bool
	S3StateBucket string
	S3StateKey    string
)

var (
	Sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	EcsClient            = ecs.New(Sess)
	ResourceGroupsClient = resourcegroupstaggingapi.New(Sess)
	S3Client             = s3.New(Sess)
)

func init() {
	Debug = getenv("DEBUG", true).(bool)
	S3StateBucket = getenv("AWS_SWITCH_STATE_BUCKET", "").(string)
	S3StateKey = getenv("AWS_SWITCH_STATE_KEY", "aws-switch").(string)
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
