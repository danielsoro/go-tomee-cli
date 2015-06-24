package execution

import (
	"runtime"

	"github.com/danielsoro/tomee-cli/execution/unix"
	"github.com/danielsoro/tomee-cli/execution/windows"
)

type Execution interface {
	Start(tomeePath string) error
	Stop(tomeePath string) error
	Restart(tomeePath string)
}

func CreateExecution() Execution {
	switch runtime.GOOS {
	case "windows":
		return new(windows.Windows)
	default:
		return new(unix.Unix)
	}
}
