package until

import (
	"os"
	"os/exec"
)

func Execute(command *exec.Cmd) {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if command.Run() != nil {
		os.Exit(1)
	}
}
