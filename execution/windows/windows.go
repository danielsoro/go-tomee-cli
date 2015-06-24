package windows

import (
	"fmt"
	"os/exec"
	"path"
)

type Windows struct {
}

func (w Windows) Start(tomeePath string) error {
	cmd := exec.Command(path.Join(tomeePath, "bin", "startup.bat"))
	return cmd.Run()
}

func (w Windows) Stop(tomeePath string) error {
	cmd := exec.Command(path.Join(tomeePath, "bin", "shutdown.bat"))
	return cmd.Run()
}

func (w Windows) Restart(tomeePath string) {
	err := w.Stop(tomeePath)
	if err != nil {
		fmt.Println("TomEE isn't started...")
	}

	fmt.Printf("Starting server..")
	err = w.Start(tomeePath)
	if err != nil {
		fmt.Println("Error during start")
		return
	}

	fmt.Println("TomEE started")
}
