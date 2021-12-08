package types

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/chrispruitt/aws-switch/config"
	log "github.com/sirupsen/logrus"
)

var RDSClusterType = "rds:cluster"

type RDSCluster struct {
	ARN string `json:"arn"`
}

func (s RDSCluster) Halt() error {
	log.Infof("Halting RDSCluster")
	_, err := config.RDSClient.StopDBCluster(&rds.StopDBClusterInput{
		DBClusterIdentifier: &s.ARN,
	})
	return err
}

func (s RDSCluster) Resume() error {
	log.Infof("Resuming RDSCluster")
	_, err := config.RDSClient.StartDBCluster(&rds.StartDBClusterInput{
		DBClusterIdentifier: &s.ARN,
	})
	return err
}

func (s RDSCluster) GetARN() string {
	return s.ARN
}

func (s RDSCluster) GetType() string {
	return RDSClusterType
}
