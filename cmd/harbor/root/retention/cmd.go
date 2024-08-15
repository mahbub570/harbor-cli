package retention

import "github.com/spf13/cobra"

func Retention() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retention",
		Short: "Manage retention policies",
		Long:  `Manage retention policies in Harbor context`,
	}
	cmd.AddCommand(
		newCreateCmd(),
		newGetCmd(),
		newUpdateCmd(),
		newDeleteCmd(),
		newTriggerCmd(),
		newListExecutionsCmd(),
		newGetTaskLogCmd(),
		newGetMetadataCmd(),
	)

	return cmd
}