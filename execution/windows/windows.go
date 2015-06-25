package windows

import (
	"os/exec"
	"path"

	"github.com/danielsoro/tomee-cli/execution"
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
	execution.Restart(w, tomeePath)
}
