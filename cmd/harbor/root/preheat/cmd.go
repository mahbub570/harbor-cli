package preheat

import (
	"github.com/spf13/cobra"
)

func Preheat() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "preheat",
		Short:   "Manage registries",
		Long:    `Manage registries in Harbor`,
		Example: `  harbor registry list`,
	}
	cmd.AddCommand(
		newCreateCommand(),
	)

	return cmd
}
