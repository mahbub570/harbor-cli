package systemcveallowlist

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/spf13/cobra"
    "github.com/goharbor/harbor-cli/pkg/utils"
    "github.com/goharbor/go-client/pkg/sdk/v2.0/client/system_cve_allowlist"
)

func getCVEAllowlistCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "get",
        Short: "Get system CVE allowlist",
        RunE: func(cmd *cobra.Command, args []string) error {
            client, err := utils.GetClient()
            if err != nil {
                return fmt.Errorf("failed to get client: %w", err)
            }

            params := system_cve_allowlist.NewGetSystemCVEAllowlistParams()
            resp, err := client.SystemCVEAllowlist.GetSystemCVEAllowlist(context.Background(), params)
            if err != nil {
                return fmt.Errorf("failed to get system CVE allowlist: %w", err)
            }
            if len(resp.Payload.Items) == 0 {
                fmt.Println("Warning: The CVE allowlist is currently empty.")
            } else {
                fmt.Printf("CVE Allowlist contains %d item(s).\n", len(resp.Payload.Items))
            }
 
            jsonData, err := json.MarshalIndent(resp.Payload, "", "  ")
            if err != nil {
                return fmt.Errorf("failed to marshal CVE allowlist: %w", err)
            }

            fmt.Println(string(jsonData))
            return nil
        },
    }
}