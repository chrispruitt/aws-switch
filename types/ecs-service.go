package types

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/chrispruitt/aws-switch/config"
	log "github.com/sirupsen/logrus"
)

var ECSServiceType = "ecs:service"

type ECSService struct {
	ARN          string `json:"arn"`
	Cluster      string `json:"cluster"`
	DesiredCount int64  `json:"desiredCount"`
}

func (s ECSService) Halt() error {
	log.Infof("Halting ECSService %v", s)
	_, err := config.EcsClient.UpdateService(&ecs.UpdateServiceInput{
		Cluster:      aws.String(s.Cluster),
		Service:      aws.String(s.ARN),
		DesiredCount: aws.Int64(0),
	})

	return err
}

func (s ECSService) Resume() error {
	log.Infof("Resuming ECSService %v", s)
	_, err := config.EcsClient.UpdateService(&ecs.UpdateServiceInput{
		Cluster:      aws.String(s.Cluster),
		Service:      aws.String(s.ARN),
		DesiredCount: aws.Int64(s.DesiredCount),
	})
	return err
}

func (s ECSService) GetARN() string {
	return s.ARN
}

func (s ECSService) GetType() string {
	return ECSServiceType
}
