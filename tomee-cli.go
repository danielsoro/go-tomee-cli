package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func start(path string) {
	cmd := exec.Command("sh", "-c", path+"/bin/startup.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func stop(path string) {
	cmd := exec.Command("sh", "-c", path+"/bin/shutdown.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func restart(path string) {
	stop(path)
	start(path)
}

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
			start(c.String("path"))
		},
	}

	stopCommand := cli.Command{
		Name:  "stop",
		Usage: "stop the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			stop(c.String("path"))
		},
	}

	restartCommand := cli.Command{
		Name:  "restart",
		Usage: "restart the TomEE server",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			restart(c.String("path"))
		},
	}

	return []cli.Command{startCommand, stopCommand, restartCommand}
}

func main() {
	app := cli.NewApp()
	app.Name = "tomee-cli"
	app.Version = "1.0.0"
	app.Commands = createCommands()
	app.Run(os.Args)
}
