package cmd

import (
	"alexejk.io/portal/pkg/config"
	"alexejk.io/portal/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RootCmd(v *version.Info) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "portal <command>",
		Short: "",
	}

	cmd.AddCommand(ConnectCmd())
	cmd.AddCommand(ListCmd())
	cmd.AddCommand(ZshCompletionCmd())

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		config.Initialize()
		initializeLogger()

		return nil
	}

	cmd.Version = v.String()
	cmd.SilenceErrors = true

	return cmd
}

func initializeLogger() {

	c, _ := config.GetConfig()

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       false,
	})

	if c.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		if level, err := logrus.ParseLevel(c.LogLevel); err == nil {
			logrus.SetLevel(level)
		}
	}
}
