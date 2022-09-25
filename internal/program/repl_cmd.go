package program

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/central182/odie/internal/adapter/inbound/cli"
	dependency_cli "github.com/central182/odie/internal/program/dependency/cli"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Look words up in an interactive session",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var config Config
		err := viper.Unmarshal(&config)
		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Print("> ")
		scn := bufio.NewScanner(os.Stdin)
		for scn.Scan() {
			headwordRaw := scn.Text()

			p := cli.NewPrinter(dependency_cli.InitInitApplication(config.AppId, config.AppKey))

			es, err := p.PrintEntriesOfHeadword(headwordRaw)
			fmt.Printf(es)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Print("> ")
		}
	},
}

func init() {
	rootCmd.AddCommand(replCmd)
}
