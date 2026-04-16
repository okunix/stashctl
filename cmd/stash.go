package cmd

import (
	"errors"
	"fmt"

	"github.com/okunix/stash-sdk/stash/v1"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:     "stash COMMAND",
	Aliases: []string{"stashes"},
	Short:   "manage stashes",
}

func init() {
	rootCmd.AddCommand(stashCmd)
	stashCmd.AddCommand(getStashByIDCmd)
	stashCmd.AddCommand(lockStashCmd)
	stashCmd.AddCommand(unlockStashCmd)
	stashCmd.AddCommand(deleteStashCmd)
	stashCmd.AddCommand(listStashesCmd)

	stashCmd.AddCommand(createStashCmd)
	createStashCmd.Flags().StringP("description", "d", "", "description of a stash")

	stashCmd.AddCommand(updateStashCmd)
	updateStashCmd.Flags().StringP("name", "n", "", "name of a stash")
	updateStashCmd.Flags().StringP("description", "d", "", "description of a stash")
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
		return printStashResponse(stash)
	},
}

var listStashesCmd = &cobra.Command{
	Use:     "list [member|maintainer]",
	Aliases: []string{"ls"},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if len(args) == 0 {
			return nil
		}
		if args[0] != "member" && args[0] != "maintainer" {
			return errors.New("invalid list argument see help")
		}
		return nil
	},
	Short: "list stashes",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashes, err := client.ListStashes(ctx)
		if err != nil {
			return err
		}
		if len(args) == 0 {
			if len(stashes.Maintainer) > 0 {
				fmt.Printf("Maintainer\n----------\n")
				printStashResponseList(stashes.Maintainer)
			}
			fmt.Println()
			if len(stashes.Member) > 0 {
				fmt.Printf("Member\n------\n")
				printStashResponseList(stashes.Member)
			}
			return nil
		}
		if args[0] == "member" {
			printStashResponseList(stashes.Member)
			return nil
		}

		if args[0] == "maintainer" {
			printStashResponseList(stashes.Maintainer)
			return nil
		}
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

var deleteStashCmd = &cobra.Command{
	Use:     "delete ID",
	Aliases: []string{"rm"},
	Short:   "delete stash",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]
		return client.DeleteStash(ctx, stashID)
	},
}

var updateStashCmd = &cobra.Command{
	Use:   "update ID [-n NAME] [-d DESCRIPTION]",
	Short: "update stash",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		stashID := args[0]

		var name *string
		nameStr, _ := cmd.Flags().GetString("name")
		if nameStr != "" {
			name = &nameStr
		}

		var desc *string
		descStr, _ := cmd.Flags().GetString("description")
		if descStr != "" {
			desc = &descStr
		}

		req := stash.UpdateStashRequest{Name: name, Description: desc}
		return client.UpdateStash(ctx, stashID, req)
	},
}
