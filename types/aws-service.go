package types

type AWSService interface {
	Halt() error
	Resume() error
	GetARN() string
	GetType() string
}
