package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type zshCompletionOpts struct {
}

func ZshCompletionCmd() *cobra.Command {

	o := &zshCompletionOpts{}

	cmd := &cobra.Command{
		Use:    "zsh-completion",
		Hidden: true,
		RunE:   o.RunE,
	}

	return cmd
}

func (o *zshCompletionOpts) RunE(cmd *cobra.Command, args []string) error {

	return cmd.GenZshCompletion(os.Stdout)
}
