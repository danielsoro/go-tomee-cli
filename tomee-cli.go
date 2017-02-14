package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/danielsoro/go-tomee-cli/deployment"
	"github.com/danielsoro/go-tomee-cli/factory"
	"github.com/danielsoro/go-tomee-cli/install"
)

func createCommands() []cli.Command {
	execution := factory.ExecutionFactory(runtime.GOOS)

	pathFlag := cli.StringFlag{
		Name:   "path",
		Usage:  "path of the TomEE server. Default value: $TOMEE_HOME",
		EnvVar: "TOMEE_HOME",
	}

	profileFlag := cli.StringFlag{
		Name:  "profile",
		Usage: "profile for the TomEE server.",
	}

	versionFlag := cli.StringFlag{
		Name:  "version",
		Usage: "version for the TomEE server.",
	}

	startCommand := cli.Command{
		Name:  "start",
		Usage: "start the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			execution.Start(c.String("path"))
		},
	}

	stopCommand := cli.Command{
		Name:  "stop",
		Usage: "stop the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			execution.Stop(c.String("path"))
		},
	}

	restartCommand := cli.Command{
		Name:  "restart",
		Usage: "restart the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			execution.Restart(c.String("path"))
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

	installCommand := cli.Command{
		Name:  "install",
		Usage: "install a version of TomEE profile",
		Flags: []cli.Flag{pathFlag, profileFlag, versionFlag},
		Action: func(c *cli.Context) {
			err := install.Install(c.String("path"), c.String("profile"), c.String("version"))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Installed in: " + c.String("path"))
		},
	}

	return []cli.Command{startCommand, stopCommand, restartCommand, deployCommand, undeployCommand, installCommand}
}

func main() {
	app := cli.NewApp()
	app.Name = "tomee-cli"
	app.Usage = "Command line tool helps system administrators and developers to manage a instance of TomEE server."
	app.Version = "1.0.0"
	app.Commands = createCommands()
	app.Run(os.Args)
}
