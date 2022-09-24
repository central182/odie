package program

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/central182/odie/internal/adapter/inbound/cli"
	dependency_cli "github.com/central182/odie/internal/program/dependency/cli"
)

var rootCmd = &cobra.Command{
	Use:  "odie HEADWORD",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var config Config
		err := viper.Unmarshal(&config)
		if err != nil {
			cobra.CheckErr(err)
		}

		p := cli.NewPrinter(dependency_cli.InitInitApplication(config.AppId, config.AppKey))

		es, err := p.PrintEntriesOfHeadword(args[0])
		fmt.Printf(es)
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("appId", "", "app_id provided by Oxford Dictionaries API")
	rootCmd.PersistentFlags().String("appKey", "", "app_key provided by Oxford Dictionaries API")
	viper.BindPFlag("appId", rootCmd.PersistentFlags().Lookup("appId"))
	viper.BindPFlag("appKey", rootCmd.PersistentFlags().Lookup("appKey"))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".odie")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}
