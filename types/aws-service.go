package types

type AWSService interface {
	Halt() error
	Start() error
	GetARN() string
	GetType() string
}
