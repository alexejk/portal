package cmd

import (
	"github.com/alexejk/portal/pkg/portal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type connectOpts struct {

}

func ConnectCmd() *cobra.Command {

	o := &connectOpts{}

	cmd := &cobra.Command{
		Use: "connect <destination>",
		RunE: o.RunE,
	}

	return cmd
}


func (o *connectOpts) RunE(cmd *cobra.Command, args []string) error {

	logrus.Info("Attempting connection")
	err := portal.Connect()
	return err
}