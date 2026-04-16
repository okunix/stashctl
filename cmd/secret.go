package cmd

import (
	"fmt"
	"time"

	"github.com/okunix/stash-sdk/stash/v1"
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:     "secret COMMAND",
	Aliases: []string{"secrets"},
	Short:   "manage secrets of a stash",
}

func init() {
	rootCmd.AddCommand(secretCmd)
	secretCmd.AddCommand(getSecretCmd)
	secretCmd.AddCommand(listSecretsCmd)
	secretCmd.AddCommand(addSecretCmd)
	secretCmd.AddCommand(deleteSecretCmd)
}

var getSecretCmd = &cobra.Command{
	Use:   "get STASH_ID NAME",
	Short: "get secret from a stash",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		name := args[1]

		secret, err := client.GetSecretsEntry(ctx, stashID, name)
		if err != nil {
			return err
		}
		fmt.Printf("%s", secret)
		return nil
	},
}

var listSecretsCmd = &cobra.Command{
	Use:     "list STASH_ID",
	Short:   "list available secrets",
	Aliases: []string{"ls"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		resp, err := client.GetSecrets(ctx, stashID)
		if err != nil {
			return err
		}
		fmt.Printf("UnlockedAt:\t%s\n\n", resp.UnlockedAt.Format(time.RFC3339))
		fmt.Printf("Secrets\n-------\n")
		for _, v := range resp.Keys {
			fmt.Println(v)
		}
		return nil
	},
}

var addSecretCmd = &cobra.Command{
	Use:     "add STASH_ID NAME VALUE",
	Short:   "add new secret to the stash",
	Aliases: []string{"create"},
	Args:    cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		name := args[1]
		value := args[2]
		req := stash.AddSecretRequest{Name: name, Value: value}
		return client.AddSecretsEntry(ctx, stashID, req)
	},
}

var deleteSecretCmd = &cobra.Command{
	Use:     "delete STASH_ID NAME",
	Short:   "delete secret from stash",
	Aliases: []string{"rm", "remove"},
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		name := args[1]
		return client.RemoveSecretsEntry(ctx, stashID, name)
	},
}
