package cmd

import (
	"context"
	"errors"

	"github.com/okunix/stash-sdk/stash/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "stashctl",
	Short: "stash secrets manager cli tool",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		serverAddr := viper.GetString("server")
		username := viper.GetString("username")
		password := viper.GetString("password")

		var client *stash.Client
		var err error

		if username != "" || password != "" {
			client, err = stash.NewClient(
				stash.WithAddr(serverAddr),
				stash.WithUser(username, password),
			)
		} else {
			client, err = stash.NewClient(stash.WithAddr(serverAddr))
		}

		if err != nil {
			return err
		}

		if err := client.Ping(ctx); err != nil {
			return errors.New("failed to reach server")
		}
		cmd.SetContext(context.WithValue(cmd.Context(), "stash-client", client))
		return nil
	},
}

var cfgFile string

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "set different config file")
}

func Execute() {
	cobra.OnInitialize(func() {
		viper.SetEnvPrefix("STASH")
		viper.AutomaticEnv()
		viper.SetDefault("server", "http://localhost:7878")
		viper.SetDefault("version", "v1")

		if cfgFile != "" {
			viper.AddConfigPath(cfgFile)
		} else {
			viper.SetConfigName("stashctl")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
			viper.AddConfigPath("$XDG_CONFIG_HOME/stash/")
			viper.AddConfigPath("$HOME/.config/stash/")
			viper.AddConfigPath("$HOME/.stash/")
			viper.AddConfigPath("/etc/stash/")
		}
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				errExit(err)
			}
		}
	})
	if err := rootCmd.Execute(); err != nil {
		errExit(err)
	}
}
