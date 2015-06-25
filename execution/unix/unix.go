package unix

import (
	"os/exec"
	"path"

	"github.com/danielsoro/tomee-cli/execution"
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

func (u Unix) Restart(tomeePath string) {
	execution.Restart(u, tomeePath)
}
