package types

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/chrispruitt/aws-switch/config"
	log "github.com/sirupsen/logrus"
)

var RDSInstanceType = "rds:db"

type RDSInstance struct {
	ARN string `json:"arn"`
}

func (s RDSInstance) Halt() error {
	log.Infof("Halting RDSInstance")
	_, err := config.RDSClient.StopDBInstance(&rds.StopDBInstanceInput{
		DBInstanceIdentifier: &s.ARN,
	})
	return err
}

func (s RDSInstance) Resume() error {
	log.Infof("Resuming RDSInstance")
	_, err := config.RDSClient.StartDBInstance(&rds.StartDBInstanceInput{
		DBInstanceIdentifier: &s.ARN,
	})
	return err
}

func (s RDSInstance) GetARN() string {
	return s.ARN
}

func (s RDSInstance) GetType() string {
	return RDSInstanceType
}
