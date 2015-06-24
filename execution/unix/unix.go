package unix

import (
	"fmt"
	"os/exec"
	"path"
)

type Unix struct {
}

func (u Unix) Start(tomeePath string) error {
	cmd := exec.Command("sh", "-c", path.Join(tomeePath, "bin", "startup.sh"))
	return cmd.Run()
}

func (u Unix) Stop(tomeePath string) error {
	cmd := exec.Command("sh", "-c", path.Join(tomeePath, "bin", "shutdown.sh"))
	return cmd.Run()
}

func (u Unix) Restart(path string) {
	err := u.Stop(path)
	if err != nil {
		fmt.Println("TomEE isn't started...")
	}

	fmt.Printf("Starting server..")
	err = u.Start(path)
	if err != nil {
		fmt.Println("Error during start")
		return
	}

	fmt.Println("TomEE started")
}
