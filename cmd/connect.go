package cmd

import (
	"errors"
	"fmt"

	"alexejk.io/portal/pkg/portal"
	"github.com/spf13/cobra"
)

type connectOpts struct {
}

func ConnectCmd() *cobra.Command {

	o := &connectOpts{}

	cmd := &cobra.Command{
		Use: "connect [destination]",
		Aliases: []string{
			"jump",
		},
		Short: "Connect to a known destination by its name",
		Args:  cobra.ExactArgs(1),
		RunE:  o.RunE,
	}

	return cmd
}

func (o *connectOpts) RunE(cmd *cobra.Command, args []string) error {

	destination := args[0]
	if destination == "" {
		return errors.New("destination cannot be empty")
	}

	registry := portal.NewRegistry()
	p, err := registry.GetPortal(destination)
	if err != nil {
		return err
	}

	fmt.Printf("Connecting to '%s'...\n", p.Name())
	if p.Hint() != "" {
		fmt.Printf(" -> Hint: %s \n", p.Hint())
	}

	return p.Connect()
}
