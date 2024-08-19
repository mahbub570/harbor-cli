package systemcveallowlist

import (
	"github.com/spf13/cobra"
)

func SystemcveallowlistCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "system-cve-allowlist",
		Short:   "Manage system CVE allowlist",
		Long:    `Manage the system-level CVE allowlist in Harbor`,
		Example: `  harbor system-cve-allowlist get
					harbor system-cve-allowlist update --file allowlist.json (eg: {"cve_allowlist": [ { "cve_id":  }]})
					harbor system-cve-allowlist update --cve-ids CVE-2021-12345,CVE-2021-67890`,
	}

	cmd.AddCommand(
		getCVEAllowlistCommand(),
		updateCVEAllowlistCommand(),
	)

	return cmd
}
