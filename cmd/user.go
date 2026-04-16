package cmd

import (
	"fmt"

	"github.com/okunix/stash-sdk/stash/v1"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"users"},
	Short:   "user management commands",
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(createUserCmd)
	userCmd.AddCommand(getUserCmd)
	userCmd.AddCommand(listUsersCmd)
	listUsersCmd.Flags().UintP("limit", "l", 30, "limit")
	listUsersCmd.Flags().UintP("offset", "o", 0, "offset")
}

var createUserCmd = &cobra.Command{
	Use:     "create USERNAME PASSWORD",
	Short:   "create new user",
	Aliases: []string{"add"},
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		username := args[0]
		password := args[1]
		req := stash.CreateUserRequest{
			Username: username,
			Password: password,
		}
		return client.CreateUser(ctx, req)
	},
}

var getUserCmd = &cobra.Command{
	Use:   "get [USER_ID|USERNAME]",
	Short: "get user by username or user id",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		var user *stash.UserResponse
		var err error
		if len(args) == 0 {
			user, err = client.Whoami(ctx)
		} else {
			user, err = client.GetUserByID(ctx, args[0])
		}
		if err != nil {
			return err
		}
		return printUserResponse(user)
	},
}

var listUsersCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list users on the system",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)
		limit, _ := cmd.Flags().GetUint("limit")
		offset, _ := cmd.Flags().GetUint("offset")
		resp, err := client.ListUsers(ctx, limit, offset)
		if err != nil {
			return err
		}
		fmt.Printf("Users\n-----\n")
		return printUserResponseList(resp.Result)
	},
}
