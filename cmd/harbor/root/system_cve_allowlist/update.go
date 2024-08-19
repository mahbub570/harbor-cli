package systemcveallowlist

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"

    "github.com/spf13/cobra"
    "github.com/goharbor/harbor-cli/pkg/utils"
    "github.com/goharbor/go-client/pkg/sdk/v2.0/client/system_cve_allowlist"
    "github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

func updateCVEAllowlistCommand() *cobra.Command {
    var filePath string
    var cveIDs []string

    cmd := &cobra.Command{
        Use:   "update",
        Short: "Update system CVE allowlist",
        Long:  `Update the system-level CVE allowlist with a new one provided in a JSON file or via the command line.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            client, err := utils.GetClient()
            if err != nil {
                return fmt.Errorf("failed to get client: %w", err)
            }

            var allowlist models.CVEAllowlist

            if filePath != "" {
                jsonData, err := ioutil.ReadFile(filePath)
                if err != nil {
                    return fmt.Errorf("failed to read file %s: %w", filePath, err)
                }

                var inputAllowlist struct {
                    CVEAllowlist []struct {
                        CVEID string `json:"cve_id"`
                    } `json:"cve_allowlist"`
                }

                if err := json.Unmarshal(jsonData, &inputAllowlist); err != nil {
                    return fmt.Errorf("failed to parse JSON: %w", err)
                }

                for _, item := range inputAllowlist.CVEAllowlist {
                    allowlist.Items = append(allowlist.Items, &models.CVEAllowlistItem{
                        CVEID: item.CVEID,
                    })
                }
            }
            if len(cveIDs) > 0 {
                for _, cveID := range cveIDs {
                    allowlist.Items = append(allowlist.Items, &models.CVEAllowlistItem{
                        CVEID: cveID,
                    })
                }
            }

            if len(allowlist.Items) == 0 {
                return fmt.Errorf("no CVE IDs provided; use either --file or --cve-ids")
            }

            params := system_cve_allowlist.NewPutSystemCVEAllowlistParams()
            params.Allowlist = &allowlist

            resp, err := client.SystemCVEAllowlist.PutSystemCVEAllowlist(context.Background(), params)
            if err != nil {
                return fmt.Errorf("failed to update system CVE allowlist: %w", err)
            }

            fmt.Println("System CVE allowlist updated successfully:", resp)
            return nil
        },
    }

    cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the JSON file containing the new CVE allowlist")
    cmd.Flags().StringSliceVarP(&cveIDs, "cve-ids", "c", []string{}, "List of CVE IDs to add to the allowlist")
    return cmd
}
