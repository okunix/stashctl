package cmd

import (
	"fmt"

	"github.com/okunix/stash-sdk/stash/v1"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:   "stash COMMAND",
	Short: "manage stashes",
}

func init() {
	rootCmd.AddCommand(stashCmd)
	stashCmd.AddCommand(getStashByIDCmd)
	stashCmd.AddCommand(lockStashCmd)
	stashCmd.AddCommand(unlockStashCmd)

	stashCmd.AddCommand(createStashCmd)
	createStashCmd.Flags().StringP("description", "d", "", "description of a stash")
}

var getStashByIDCmd = &cobra.Command{
	Use:   "get ID",
	Short: "get stash by id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stash, err := client.GetStashByID(ctx, args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", stash)
		return nil
	},
}

var unlockStashCmd = &cobra.Command{
	Use:   "unlock ID PASSWORD",
	Short: "unlock stash",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		password := args[1] // add more secure approach
		return client.Unlock(ctx, stashID, password)
	},
}

var lockStashCmd = &cobra.Command{
	Use:   "lock ID",
	Short: "lock stash",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		return client.Lock(ctx, stashID)
	},
}

var createStashCmd = &cobra.Command{
	Use:   "create NAME PASSWORD [-d DESCRIPTION]",
	Short: "create stash",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		name := args[0]
		password := args[1]
		var description *string
		descriptionStr, err := cmd.Flags().GetString("description")
		if err == nil {
			description = &descriptionStr
		}
		client := mustGetStashClient(ctx)
		return client.CreateStash(ctx, stash.CreateStashRequest{
			Name:        name,
			Password:    password,
			Description: description,
		})
	},
}
