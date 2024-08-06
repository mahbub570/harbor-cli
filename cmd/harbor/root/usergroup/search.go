package usergroup

import (
    "bufio"
    "fmt"
    "os"
    "strings"
 

    "github.com/goharbor/harbor-cli/pkg/api"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

func UserGroupsSearchCmd() *cobra.Command {
    var groupName string

    cmd := &cobra.Command{
        Use:   "search [groupName]",
        Short: "search user groups",
        Args:  cobra.MaximumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) > 0 {
                groupName = args[0]
            }

            if groupName == "" {
                fmt.Print("Enter group name: ")
                reader := bufio.NewReader(os.Stdin)
                input, _ := reader.ReadString('\n')
                groupName = strings.TrimSpace(input)
            }
            // Clear the previous message
            fmt.Print("\033[K") // ANSI escape code to clear the line

            fmt.Printf("Searching for groups with name '%s'...\r", groupName)
            response, err := api.SearchUserGroups(groupName)
            if err != nil {
                log.Errorf("failed to search user groups: %v", err)
                return
            }

           
            if len(response.Payload) == 0 {
                log.Infof("No user groups found with the name %s", groupName)
                return
            }

            for _, group := range response.Payload {
                log.Infof("ID: %d, Name: %s, Type: %d", group.ID, group.GroupName, group.GroupType)
            }
 
        },
    }

    flags := cmd.Flags()
    flags.StringVarP(&groupName, "name", "n", "", "Group name to search")

    return cmd
}
