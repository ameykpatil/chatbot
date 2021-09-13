package main

import (
	"context"
	"log"
	"os"

	"github.com/ameykpatil/chatbot/command"
	"github.com/ameykpatil/chatbot/config"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "chatbot"
	app.Usage = "The service to serve the replies for the messages based on intent"
	app.UsageText = "chatbot"

	// Init envConfig
	envConfig, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("envConfig could not be loaded, error : %s", err)
	}

	baseCommand := command.NewBaseCommand(context.TODO(), envConfig)
	httpCommand := command.NewHTTPCommand(baseCommand)

	app.Commands = []cli.Command{
		{
			Name:   "http",
			Usage:  "Start REST API service",
			Action: httpCommand.Run,
		},
		// Add more commands such as consumer or some periodic job
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
