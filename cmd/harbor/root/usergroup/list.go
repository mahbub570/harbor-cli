package usergroup

import (
    "github.com/goharbor/harbor-cli/pkg/api"
    "github.com/goharbor/harbor-cli/pkg/utils"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

func UserGroupsListCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "list",
        Short: "list user groups",
        Args:  cobra.NoArgs,
        Run: func(cmd *cobra.Command, args []string) {
            response, err := api.ListUserGroups()
            if err != nil {
                log.Errorf("failed to list user groups: %v", err)
                return
            }

            FormatFlag := viper.GetString("output-format")
            if FormatFlag != "" {
                utils.PrintPayloadInJSONFormat(response.Payload)
            } else {
                for _, group := range response.Payload {
                    log.Infof("ID: %d, Name: %s, Type: %d", group.ID, group.GroupName, group.GroupType)
                }
            }
        },
    }

    return cmd
}
