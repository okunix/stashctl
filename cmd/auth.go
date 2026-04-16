package cmd

import (
	"fmt"
	"syscall"

	"github.com/okunix/stash-sdk/stash/v1"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var authCmd = &cobra.Command{
	Use:   "auth COMMAND",
	Short: "auth related commands",
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(changePasswordCmd)
	authCmd.AddCommand(tokenCmd)
}

var changePasswordCmd = &cobra.Command{
	Use:     "change-password [OLD_PASSWORD] [NEW_PASSWORD]",
	Aliases: []string{"passwd"},
	Args:    cobra.MaximumNArgs(2),
	Short:   "change password of a current user",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)

		var oldPassword, newPassword string
		var err error
		switch len(args) {
		case 2:
			oldPassword = args[0]
			newPassword = args[1]
		case 1:
			oldPassword = args[0]
			fmt.Printf("New password: ")
			newPasswordBytes, _ := term.ReadPassword(syscall.Stdin)
			newPassword = string(newPasswordBytes)
			fmt.Println()
		default:
			fmt.Printf("Old password: ")
			oldPasswordBytes, _ := term.ReadPassword(syscall.Stdin)
			oldPassword = string(oldPasswordBytes)
			fmt.Printf("\nNew password: ")
			newPasswordBytes, _ := term.ReadPassword(syscall.Stdin)
			newPassword = string(newPasswordBytes)
			fmt.Println()
		}
		if err != nil {
			return err
		}

		req := stash.ChangePasswordRequest{
			OldPassword: oldPassword,
			NewPassword: newPassword,
		}
		if err := client.ChangePassword(ctx, req); err != nil {
			return err
		}
		fmt.Println("password changed")
		return nil
	},
}

var tokenCmd = &cobra.Command{
	Use:   "token USERNAME [PASSWORD]",
	Args:  cobra.RangeArgs(1, 2),
	Short: "get jwt token",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		client := mustGetStashClient(ctx)

		username := args[0]
		var password string
		if len(args) == 1 {
			fmt.Scanf("%s", &password)
		} else if len(args) >= 2 {
			password = args[1]
		}

		req := stash.GetTokenRequest{Username: username, Password: password}
		resp, err := client.GetToken(ctx, req)
		if err != nil {
			return err
		}
		fmt.Printf("%s", resp.Token)
		return nil
	},
}
