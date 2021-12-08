package lib

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/chrispruitt/aws-switch/config"
	. "github.com/chrispruitt/aws-switch/types"
)

func GetResourceArns(tags map[string]string, resourceType string) ([]string, error) {

	tagFilters := []*resourcegroupstaggingapi.TagFilter{}

	for k, v := range tags {
		tagFilters = append(tagFilters, &resourcegroupstaggingapi.TagFilter{
			Key:    aws.String(k),
			Values: aws.StringSlice(strings.Split(v, ",")),
		})
	}

	arns := []string{}
	input := &resourcegroupstaggingapi.GetResourcesInput{
		ResourceTypeFilters: aws.StringSlice([]string{resourceType}),
		TagFilters:          tagFilters,
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

	// Get ECS Services
	ecsServices, err := GetECSServices(tags)
	if err != nil {
		return nil, err
	}
	services = append(services, ecsServices...)

	// Get RDS Clusters
	rdsClusters, err := GetRDSClusters(tags)
	if err != nil {
		return nil, err
	}
	services = append(services, rdsClusters...)

	// Get RDS Clusters
	rdsInstances, err := GetRDSInstances(tags)
	if err != nil {
		return nil, err
	}
	services = append(services, rdsInstances...)

	// TODO EMR
	// TODO etc...

	return services, nil
}
