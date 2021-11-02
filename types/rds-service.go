package types

import (
	log "github.com/sirupsen/logrus"
)

type RDSService struct {
	ARN string
}

func (s RDSService) Halt() error {
	log.Infof("Halting RDSService")
	return nil
}

func (s RDSService) Resume() error {
	log.Infof("Resuming RDSService")
	return nil
}

func (s RDSService) GetARN() string {
	return s.ARN
}
