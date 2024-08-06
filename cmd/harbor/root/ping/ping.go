package ping

import (
    "fmt"
    "github.com/goharbor/go-client/pkg/sdk/v2.0/client/ldap"
    "github.com/goharbor/harbor-cli/pkg/utils"
    "github.com/spf13/cobra"
    "github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

func newLDAPPingCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "ldap",
        Short: "Ping LDAP server",
        RunE: func(cmd *cobra.Command, args []string) error {
            ctx, client, err := utils.ContextWithClient()
            if err != nil {
                return fmt.Errorf("failed to create client: %v", err)
            }

            // Create a new PingLdapParams with default values
            params := ldap.NewPingLdapParams()

            // Populate LDAP configuration
            ldapConf := &models.LdapConf{}
            params.WithLdapconf(ldapConf)

            resp, err := client.Ldap.PingLdap(ctx, params)
            if err != nil {
                if internalErr, ok := err.(*ldap.PingLdapInternalServerError); ok {
                    return fmt.Errorf("failed to ping LDAP server: %v, errors: %+v", internalErr.XRequestID, internalErr.Payload)
                }
                return fmt.Errorf("failed to ping LDAP server: %v", err)
            }

            fmt.Println("LDAP server pinged successfully")
            fmt.Printf("Response: %+v\n", resp)
            return nil
        },
    }

    return cmd
}
