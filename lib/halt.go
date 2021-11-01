package lib

import (
	"fmt"

	"github.com/chrispruitt/aws-switch/state"
)

func Halt(tags map[string]string) error {
	services, err := GetAWSServices(tags)
	if err != nil {
		return err
	}

	for _, service := range services {

		serviceState, err := state.GetService(service.GetARN())
		if err != nil {
			return err
		}
		err = service.Halt()
		if err != nil {
			return fmt.Errorf("Error halting service: %s", err)
		}

		// Keep previous state if exists
		if serviceState == nil {
			err = state.PutService(service)
			if err != nil {
				return fmt.Errorf("Error saving service state: %s", err)
			}
		}
	}
	return nil
}
