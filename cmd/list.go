package cmd

import (
	"fmt"

	"alexejk.io/portal/pkg/portal"
	"github.com/spf13/cobra"
)

type listOpts struct {
}

func ListCmd() *cobra.Command {

	o := &listOpts{}

	cmd := &cobra.Command{
		Use:  "list",
		RunE: o.RunE,
	}

	return cmd
}

func (o *listOpts) RunE(cmd *cobra.Command, args []string) error {

	reg := portal.NewRegistry()
	portals := reg.ListPortals()

	for i, p := range portals {
		if i > 0 {
			fmt.Print("  ")
		}
		fmt.Print(p.Name())
	}

	fmt.Print("\n")

	return nil
}
