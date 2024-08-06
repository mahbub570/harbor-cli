package usergroup

import (
    "strconv"

    "github.com/goharbor/harbor-cli/pkg/api"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

func UserGroupUpdateCmd() *cobra.Command {
    var groupName string
    var groupType int64

    cmd := &cobra.Command{
        Use:   "update [groupID]",
        Short: "update user group",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            groupId, err := strconv.ParseInt(args[0], 10, 64)
            if err != nil {
                log.Errorf("invalid group ID: %v", err)
                return
            }

            err = api.UpdateUserGroup(groupId, groupName, groupType)
            if err != nil {
                log.Errorf("failed to update user group: %v", err)
            } else {
                log.Infof("User group `%s` updated successfully", groupName)
            }
        },
    }

    flags := cmd.Flags()
    flags.StringVarP(&groupName, "name", "n", "", "Group name")
    flags.Int64VarP(&groupType, "type", "t", 0, "Group type")
    cmd.MarkFlagRequired("name")
    cmd.MarkFlagRequired("type")

    return cmd
}
