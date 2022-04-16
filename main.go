package main

import (
	"errors"
	"fmt"
	"lambdaenvs/pkg/entrypoint"
	"lambdaenvs/pkg/services"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "lambdaenvs",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "region",
				Value: "us-east-1",
				Usage: "AWS region",
			},
			&cli.StringFlag{
				Name:  "profile",
				Value: "default",
				Usage: "Profile name",
			},
		},
		Action: func(c *cli.Context) error {
			functionName := c.Args().Get(0)
			if len(functionName) == 0 {
				return errors.New(color.Colorize(color.Red, "the functionName must be defined"))
			}
			region := c.String("region")
			profile := c.String("profile")

			service := entrypoint.GetEnvsService(profile, region)
			output, err := service.Get(services.EnvsInput{
				FunctionName: functionName,
			})

			if err != nil {
				panic(err)
			}

			for _, env := range output.Envs {
				fmt.Println(fmt.Sprintf("%s=%s", env.Name, env.Value))
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
