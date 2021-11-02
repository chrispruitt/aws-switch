package lib

import (
	"fmt"

	"github.com/chrispruitt/aws-switch/state"
	. "github.com/chrispruitt/aws-switch/types"
)

func Halt(service AWSService) error {
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
	return nil
}
