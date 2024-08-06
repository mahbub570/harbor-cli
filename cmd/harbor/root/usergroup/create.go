package usergroup

import (
    "fmt"
    "strconv"
    "github.com/goharbor/harbor-cli/pkg/api"
    "github.com/spf13/cobra"
)

func UserGroupCreatecmd() *cobra.Command {
    var groupName string
    var groupType int64
    var ldapGroupDn string

    cmd := &cobra.Command{
        Use:   "create",
        Short: "create user group",
        Args:  cobra.NoArgs,
        RunE: func(cmd *cobra.Command, args []string) error {
            if groupName == "" {
                fmt.Print("Enter group name: ")
                fmt.Scanln(&groupName)
            }

            for {
                if groupType == 0 {
                    fmt.Print("Enter group type (1 for LDAP, 2 for HTTP, 3 for OIDC group): ")
                    var input string
                    fmt.Scanln(&input)
                    var err error
                    groupType, err = strconv.ParseInt(input, 10, 64)
                    if err != nil {
                        fmt.Println("Invalid input, please enter an integer.")
                        groupType = 0
                        continue
                    }
                }

                if groupType < 1 || groupType > 3 {
                    fmt.Println("Invalid group type. Must be 1 (LDAP), 2 (HTTP), or 3 (OIDC).")
                    groupType = 0
                    continue
                }

                if groupType == 1 {
                    fmt.Print("Enter the DN of the LDAP group: ")
                    fmt.Scanln(&ldapGroupDn)
                }

                break
            }

            fmt.Printf("Creating user group with name: %s, type: %d, LDAP DN: %s\n", groupName, groupType, ldapGroupDn)
            err := api.CreateUserGroup(groupName, groupType, ldapGroupDn)
            if err != nil {
                return fmt.Errorf("failed to create user group: %v", err)
            }
            
            fmt.Printf("User group '%s' created successfully\n", groupName)
            return nil
        },
    }

    flags := cmd.Flags()
    flags.StringVarP(&groupName, "name", "n", "", "Group name")
    flags.Int64VarP(&groupType, "type", "t", 0, "Group type")
    flags.StringVarP(&ldapGroupDn, "ldap-dn", "l", "", "The DN of the LDAP group if group type is 1 (LDAP group)")

    return cmd
}