package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func start(Path string) {
	cmd := exec.Command("sh", "-c", Path+"/bin/startup.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func stop(Path string) {
	cmd := exec.Command("sh", "-c", Path+"/bin/shutdown.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func restart(Path string) {
	stop(Path)
	start(Path)
}

func createCommands() []cli.Command {
	startCommand := cli.Command{
		Name:  "start",
		Usage: "start the TomEE server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "path",
				Usage:  "path of the TomEE server. Default value: $TOMEE_HOME",
				EnvVar: "TOMEE_HOME",
			},
		},
		Action: func(c *cli.Context) {
			start(c.String("path"))
		},
	}
	return []cli.Command{startCommand}
}

func main() {
	app := cli.NewApp()
	app.Name = "tomee-cli"
	app.Version = "1.0.0"
	app.Commands = createCommands()
	app.Run(os.Args)
}
