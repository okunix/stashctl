package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "manage stashes",
}

func init() {
	rootCmd.AddCommand(stashCmd)
	stashCmd.AddCommand(getStashByIDCmd)
	stashCmd.AddCommand(lockStash)
	stashCmd.AddCommand(unlockStash)
}

var getStashByIDCmd = &cobra.Command{
	Use:   "get",
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

var unlockStash = &cobra.Command{
	Use:   "unlock",
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

var lockStash = &cobra.Command{
	Use:   "lock",
	Short: "lock stash",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		return client.Lock(ctx, stashID)
	},
}
