package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/termie/go-shutil"
)

func start(path string) error {
	cmd := exec.Command("sh", "-c", path+"/bin/startup.sh")
	return cmd.Run()
}

func stop(path string) error {
	cmd := exec.Command("sh", "-c", path+"/bin/shutdown.sh")
	return cmd.Run()
}

func restart(path string) {
	err := stop(path)
	if err != nil {
		fmt.Println("TomEE isn't started...")
	}

	fmt.Printf("Starting server..")
	err = start(path)
	if err != nil {
		fmt.Println("Error during start")
		return
	}

	fmt.Println("TomEE started")
}

func deploy(tomeePath, packageForDeploy string) error {

	_, packageName := path.Split(packageForDeploy)
	deployPath := tomeePath + "/webapps/"
	if strings.HasSuffix(packageForDeploy, ".ear") {
		deployPath = tomeePath + "/apps/"
		os.Mkdir(deployPath, 0744)
	}

	err := shutil.CopyFile(packageForDeploy, deployPath+packageName, true)
	if err != nil {
		return err
	}

	return nil
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

	deployCommand := cli.Command{
		Name:  "deploy",
		Usage: "deploy war/ear in TomEE",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			err := deploy(c.String("path"), c.Args().First())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Deployed in: " + c.String("path"))
		},
	}

	return []cli.Command{startCommand, stopCommand, restartCommand, deployCommand}
}

func main() {
	app := cli.NewApp()
	app.Name = "tomee-cli"
	app.Version = "1.0.0"
	app.Commands = createCommands()
	app.Run(os.Args)
}
