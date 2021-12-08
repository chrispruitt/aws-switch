package lib

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/chrispruitt/aws-switch/config"
	"github.com/chrispruitt/aws-switch/types"
)

func GetRDSInstances(tags map[string]string) ([]types.AWSService, error) {
	RDSInstances := []types.AWSService{}
	RDSInstanceArns, err := GetResourceArns(tags, types.RDSInstanceType)
	if err != nil {
		return nil, fmt.Errorf("Error getting service arns: %s", err)
	}

	output, err := config.RDSClient.DescribeDBInstances(&rds.DescribeDBInstancesInput{
		Filters: []*rds.Filter{
			{
				Name:   aws.String("db-instance-id"),
				Values: aws.StringSlice(RDSInstanceArns),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Error describing service: %s", err)
	}
	for _, instance := range output.DBInstances {
		// Filter out instances that are a part of a DB Cluster - they will be handled by the RDSCluster type
		if instance.DBClusterIdentifier == nil {
			RDSInstances = append(RDSInstances, types.RDSInstance{
				ARN: *instance.DBInstanceArn,
			})
		}
	}

	return RDSInstances, nil
}
