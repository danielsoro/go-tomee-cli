package factory

import (
	"runtime"

	"github.com/danielsoro/tomee-cli/execution"
	"github.com/danielsoro/tomee-cli/execution/unix"
	"github.com/danielsoro/tomee-cli/execution/windows"
)

func ExecutionFactory(os string) execution.Execution {
	switch runtime.GOOS {
	case "windows":
		return new(windows.Windows)
	default:
		return new(unix.Unix)
	}
}
