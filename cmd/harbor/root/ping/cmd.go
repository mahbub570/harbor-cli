package ping
 
import (
	"github.com/spf13/cobra"
)

func Ping() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ping",
		Short:   "Manage ping",
		Long:    `Manage ping in Harbor`,
		Example: `  harbor ping`,
	}

	cmd.AddCommand(
		newLDAPPingCmd(),
	)

	return cmd
}
