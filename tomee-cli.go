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

func start(tomeePath string) error {
	cmd := exec.Command("sh", "-c", path.Join(tomeePath, "bin", "startup.sh"))
	return cmd.Run()
}

func stop(tomeePath string) error {
	cmd := exec.Command("sh", "-c", path.Join(tomeePath, "bin", "shutdown.sh"))
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

func undeploy(tomeePath, packageForUndeploy string) error {
	_, packageName := path.Split(packageForUndeploy)
	deployPath := path.Join(tomeePath, "webapps")
	if strings.HasSuffix(packageForUndeploy, ".ear") {
		deployPath = path.Join(tomeePath, "apps")
	}

	err := os.RemoveAll(path.Join(deployPath, packageName))
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(deployPath, strings.TrimSuffix(packageName, path.Ext(packageName))))
	if err != nil {
		return err
	}

	return nil
}

func deploy(tomeePath, packageForDeploy string) error {
	_, packageName := path.Split(packageForDeploy)
	deployPath := path.Join(tomeePath, "webapps")
	if strings.HasSuffix(packageForDeploy, ".ear") {
		deployPath = path.Join(tomeePath, "apps")
		os.Mkdir(deployPath, 0744)
	}

	err := shutil.CopyFile(packageForDeploy, path.Join(deployPath, packageName), true)
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

	undeployCommand := cli.Command{
		Name:  "undeploy",
		Usage: "undeploy war/ear in TomEE",
		Flags: []cli.Flag{pathFlag},
		Action: func(c *cli.Context) {
			err := undeploy(c.String("path"), c.Args().First())
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
			err := deploy(c.String("path"), c.Args().First())
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
