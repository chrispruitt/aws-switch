package lib

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/chrispruitt/aws-switch/config"
	"github.com/chrispruitt/aws-switch/types"
)

func GetRDSClusters(tags map[string]string) ([]types.AWSService, error) {
	RDSClusters := []types.AWSService{}
	RDSClusterArns, err := GetResourceArns(tags, types.RDSClusterType)
	if err != nil {
		return nil, fmt.Errorf("Error getting service arns: %s", err)
	}

	output, err := config.RDSClient.DescribeDBClusters(&rds.DescribeDBClustersInput{
		Filters: []*rds.Filter{
			{
				Name:   aws.String("db-cluster-id"),
				Values: aws.StringSlice(RDSClusterArns),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Error describing service: %s", err)
	}
	for _, cluster := range output.DBClusters {
		// Only pertains to provisioned rds clusters
		if *cluster.EngineMode == "provisioned" {
			RDSClusters = append(RDSClusters, types.RDSCluster{
				ARN: *cluster.DBClusterArn,
			})
		}
	}

	return RDSClusters, nil
}
