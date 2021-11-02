package state

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/chrispruitt/aws-switch/config"
	. "github.com/chrispruitt/aws-switch/types"

	log "github.com/sirupsen/logrus"
)

var (
	state = State{}
)

type UntypedJson map[string]interface{}

type State map[string]StateAWSService

type StateAWSService struct {
	Type    string
	Service json.RawMessage
}

func InitalizeState() error {
	err := state.readFromS3()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				writeErr := state.writeToS3()
				if writeErr != nil {
					return fmt.Errorf("Unable to initialize state. %v", err)
				}
			default:
				return fmt.Errorf("Unable to read state from s3. %v", err)
			}
		} else {
			return fmt.Errorf("Unable to read state from s3. %v", err)
		}
	}
	return nil
}

func GetService(arn string) (AWSService, error) {
	state.readFromS3()

	if unTypedService, ok := state[arn]; ok {
		switch unTypedService.Type {
		case ECSServiceType:
			var ecsService ECSService
			err := json.Unmarshal(unTypedService.Service, &ecsService)
			if err != nil {
				return nil, fmt.Errorf("Error unmarshaling json: %v", err)
			}
			return ecsService, nil
		default:
			log.Infof("No type found readying from state: %v", unTypedService)
		}
	}

	return nil, nil
}

func PutService(service AWSService) error {
	err := state.readFromS3()
	if err != nil {
		return fmt.Errorf("Unable to read state from s3. %v", err)
	}

	jsonService, err := json.Marshal(service)
	if err != nil {
		return err
	}

	state[service.GetARN()] = StateAWSService{
		Type:    service.GetType(),
		Service: jsonService,
	}

	state.writeToS3()

	return nil
}

func DeleteService(arn string) error {
	err := state.readFromS3()
	if err != nil {
		return fmt.Errorf("Unable to read state from s3. %v", err)
	}

	delete(state, arn)

	state.writeToS3()

	return nil
}

func (s *State) writeToS3() error {

	// Convert struct to json formated byte array
	p, err := json.Marshal(s)
	if err != nil {
		return err
	}

	// Push to s3
	putObjectInput := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(bytes.NewReader(p)),
		Bucket: aws.String(config.GetS3StateBucket()),
		Key:    aws.String(config.S3StateKey),
	}

	_, err = config.S3Client.PutObject(putObjectInput)

	if err != nil {
		return err
	}

	log.Debugf("State saved to s3://%s/%s", config.GetS3StateBucket(), config.S3StateKey)

	return nil
}

func (s *State) readFromS3() error {
	result, err := config.S3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.GetS3StateBucket()),
		Key:    aws.String(config.S3StateKey),
	})
	if err != nil {
		return err
	}
	defer result.Body.Close()
	body1, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	bodyString1 := fmt.Sprintf("%s", body1)

	decoder := json.NewDecoder(strings.NewReader(bodyString1))
	err = decoder.Decode(&s)

	if err != nil {
		return err
	}

	return nil
}
