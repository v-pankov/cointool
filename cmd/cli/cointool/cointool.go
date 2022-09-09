package cointool

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
)

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "cointool",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	command.PersistentFlags().StringVarP(
		&variable.ApiKey,
		"api-key", "k", viper.GetString("api.key"),
		"coinmarketcap API key",
	)

	command.PersistentFlags().StringVarP(
		&variable.ApiPrefix,
		"api-prefix", "p", viper.GetString("api.prefix"),
		"coinmarketcap API prefix: sandbox or pro",
	)

	command.PersistentFlags().DurationVarP(
		&variable.Timeout,
		"timeout", "t", viper.GetDuration("timeout"),
		"command timeout duration",
	)

	return command
}

func SetupConfig(configFileName, configFileLocation string) error {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	viper.SetDefault("api.key", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
	viper.SetDefault("api.prefix", "sandbox")
	viper.SetDefault("timeout", 7*time.Second)

	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFileLocation)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return fmt.Errorf("read config: %w", err)
		}
	}

	return nil
}
