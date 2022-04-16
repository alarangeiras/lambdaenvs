package entrypoint

import (
	"context"
	"lambdaenvs/pkg/services"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func GetEnvsService(profile, region string) services.EnvsService {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile))
	if err != nil {
		panic(err)
	}
	client := lambda.NewFromConfig(cfg)
	return services.NewEnvService(client)

}
