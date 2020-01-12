package portal

import (
	"fmt"
	"strings"

	"alexejk.io/portal/pkg/config"
	"alexejk.io/portal/pkg/run"
)

// RawPortal is a direct connection portal that requires no additional configuration besides a command to execute
type RawPortal struct {
	*portalShared

	cmdExecutable string
	cmdArgs       []string
}

func newRawPortal(name string, raw *config.PortalRawConfig) (*RawPortal, error) {

	commandParts := strings.Split(*raw.Command, " ")

	if len(commandParts) < 2 {
		return nil, fmt.Errorf("invalid format: expecting at least an executable and one argument")
	}

	return &RawPortal{
		portalShared: &portalShared{
			name: name,
		},

		cmdExecutable: commandParts[0],
		cmdArgs:       commandParts[1:],
	}, nil
}

func (p *RawPortal) Connect() error {

	c := run.NewRunner(p.cmdExecutable, p.cmdArgs...)

	return c.Run()
}
