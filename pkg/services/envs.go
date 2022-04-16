package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type EnvsInput struct {
	FunctionName string
}

type EnvElement struct {
	Name  string
	Value string
}

type EnvsOutput struct {
	Envs []EnvElement
}

type EnvsService interface {
	Get(input EnvsInput) (*EnvsOutput, error)
}

type envServiceImpl struct {
	client *lambda.Client
}

func NewEnvService(client *lambda.Client) EnvsService {
	return &envServiceImpl{client}
}

func (e envServiceImpl) Get(input EnvsInput) (*EnvsOutput, error) {
	output, err := e.client.GetFunction(context.Background(), &lambda.GetFunctionInput{
		FunctionName: &input.FunctionName,
	})
	if err != nil {
		return nil, err
	}
	var envsOutput EnvsOutput
	envs := envsOutput.Envs
	for key, value := range output.Configuration.Environment.Variables {
		envs = append(envs, EnvElement{
			Name:  key,
			Value: value,
		})
	}

	return &envsOutput, nil
}
