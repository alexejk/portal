package cmd

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	
	cmd := &cobra.Command{
		Use:   "portal <command>",
		Short: "",
	}

	cmd.AddCommand(ConnectCmd())

	return cmd
}
