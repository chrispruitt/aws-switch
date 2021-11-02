package lib

import (
	"fmt"

	"github.com/chrispruitt/aws-switch/state"
	. "github.com/chrispruitt/aws-switch/types"
)

func Resume(service AWSService) error {

	serviceState, err := state.GetService(service.GetARN())
	if err != nil {
		return err
	}
	// Remove service from state
	if serviceState != nil {
		err = serviceState.Resume()
		if err != nil {
			return fmt.Errorf("Error resuming service: %s", err)
		}

		err = state.DeleteService(service.GetARN())
		if err != nil {
			return fmt.Errorf("Error saving service state: %s", err)
		}
	}
	return nil
}
