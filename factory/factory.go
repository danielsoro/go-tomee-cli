package factory

import (
	"github.com/danielsoro/go-tomee-cli/execution"
	"github.com/danielsoro/go-tomee-cli/execution/unix"
	"github.com/danielsoro/go-tomee-cli/execution/windows"
)

func ExecutionFactory(os string) execution.Execution {
	switch os {
	case "windows":
		return new(windows.Windows)
	default:
		return new(unix.Unix)
	}
}
