package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/danielsoro/tomee-cli/deployment"
	"github.com/danielsoro/tomee-cli/execution"
)

func createCommands() []cli.Command {
	pathFlag := cli.StringFlag{
		Name:   "path",
		Usage:  "path of the TomEE server. Default value: $TOMEE_HOME",
		EnvVar: "TOMEE_HOME",
	}

	startCommand := cli.Command{
		Name:  "start",
		Usage: "start the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			execution.CreateExecution().Start(c.String("path"))
		},
	}

	stopCommand := cli.Command{
		Name:  "stop",
		Usage: "stop the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			execution.CreateExecution().Stop(c.String("path"))
		},
	}

	restartCommand := cli.Command{
		Name:  "restart",
		Usage: "restart the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			execution.CreateExecution().Restart(c.String("path"))
		},
	}

	undeployCommand := cli.Command{
		Name:  "undeploy",
		Usage: "undeploy war/ear in TomEE",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			err := deployment.Undeploy(c.String("path"), c.Args().First())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Undeployed in: " + c.String("path"))
		},
	}

	deployCommand := cli.Command{
		Name:  "deploy",
		Usage: "deploy war/ear in TomEE",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			err := deployment.Deploy(c.String("path"), c.Args().First())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Deployed in: " + c.String("path"))
		},
	}

	return []cli.Command{startCommand, stopCommand, restartCommand, deployCommand, undeployCommand}
}

func main() {
	app := cli.NewApp()
	app.Name = "tomee-cli"
	app.Usage = "Command line tool helps system administrators and developers to manage a instance of TomEE server."
	app.Version = "1.0.0"
	app.Commands = createCommands()
	app.Run(os.Args)
}
