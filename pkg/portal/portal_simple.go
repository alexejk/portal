package portal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// SimplePortal is a direct connection portal that requires no additional configuration besides a command to execute
type SimplePortal struct {
	*portalShared

	command string

	cmdExecutable string
	cmdArgs       []string
}

func newSimplePortal(name string, command string) (*SimplePortal, error) {

	commandParts := strings.Split(command, " ")

	if len(commandParts) < 2 {
		return nil, fmt.Errorf("invalid format: expecting at least an executable and one argument")
	}

	return &SimplePortal{
		portalShared: &portalShared{
			name: name,
		},
		command: command,

		cmdExecutable: commandParts[0],
		cmdArgs:       commandParts[1:],
	}, nil
}

func (p *SimplePortal) Connect() error {

	c := exec.Command(p.cmdExecutable, p.cmdArgs...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
