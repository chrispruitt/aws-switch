package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/chrispruitt/aws-switch/config"
	. "github.com/chrispruitt/aws-switch/types"
)

func GetResourceArns(tags map[string]string, resourceType string) ([]string, error) {
	arns := []string{}
	input := &resourcegroupstaggingapi.GetResourcesInput{
		ResourceTypeFilters: aws.StringSlice([]string{resourceType}),
	}

	pageNum := 0
	err := config.ResourceGroupsClient.GetResourcesPages(input,
		func(page *resourcegroupstaggingapi.GetResourcesOutput, lastPage bool) bool {
			pageNum++
			for _, resourceTagMapping := range page.ResourceTagMappingList {
				arns = append(arns, *resourceTagMapping.ResourceARN)
			}
			return pageNum <= 10
		})
	if err != nil {
		return nil, err
	}

	return arns, nil
}

func GetAWSServices(tags map[string]string) ([]AWSService, error) {
	services := []AWSService{}
	ecsServices, err := GetECSServices(tags)
	if err != nil {
		return nil, err
	}
	services = append(services, ecsServices...)

	// TODO RDS
	// TODO EMR
	// TODO etc...

	return services, nil
}
