package lib

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/chrispruitt/aws-switch/config"
	. "github.com/chrispruitt/aws-switch/types"
)

func GetECSServices(tags map[string]string) ([]AWSService, error) {
	ecsServices := []AWSService{}
	ecsServiceArns, err := GetResourceArns(tags, "ecs:service")
	if err != nil {
		return nil, fmt.Errorf("Error getting service arns: %s", err)
	}

	for _, arn := range ecsServiceArns {
		cluster := getClusterName(arn)
		output, err := config.EcsClient.DescribeServices(&ecs.DescribeServicesInput{
			Cluster:  aws.String(cluster),
			Services: aws.StringSlice([]string{arn}),
		})
		if err != nil {
			return nil, fmt.Errorf("Error describing service: %s", err)
		}
		for _, service := range output.Services {
			ecsServices = append(ecsServices, ECSService{
				ARN:          *service.ServiceArn,
				Cluster:      cluster,
				DesiredCount: *service.DesiredCount,
			})
		}
	}

	return ecsServices, nil
}

func getClusterName(arn string) string {
	re := regexp.MustCompile("service/(.*?)/")
	return re.FindStringSubmatch(arn)[1]
}
