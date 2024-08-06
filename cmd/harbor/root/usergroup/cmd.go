package usergroup
 
import (
	"github.com/spf13/cobra"
)

func Usergroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "usergroup",
		Short:   "Manage usergroup",
		Long:    `Manage usergroup in Harbor`,
		Example: `  harbor usergroup list`,
	}

	cmd.AddCommand(
		UserGroupCreatecmd(),
		UserGroupsListCmd(),
		UserGroupDeleteCmd(),
		UserGroupsSearchCmd(),
		UserGroupUpdateCmd(),
		UserGroupGetCmd(),

	)

	return cmd
}
