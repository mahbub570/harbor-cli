package retention

import (
	"github.com/spf13/cobra"
)

func Retention() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retention",
		Short: "Manage retention policies in Harbor",
		Long:  `Manage retention policies, executions, and related resources in Harbor`,
	}

	cmd.AddCommand(
		CreateRetentionCommand(),
		DeleteRetentionCommand(),
		ListRetentionCommand(),
		UpdateRetentionCommand(),
		// GetRetentionCommand(),
		// ListRetentionCommand(),
		// TriggerRetentionCommand(),
	)

	return cmd
}
