package main

import (
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/cmd/cli/commands"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/spf13/cobra"
	"os"
)

const (
	cliName = "lnkctl"
)

func main() {
	var command *cobra.Command

	command = NewCommand()

	err := command.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewCommand() *cobra.Command {

	var (
		clientOpts apiclient.ApiClientOpts
	)

	var command = &cobra.Command{
		Use:   cliName,
		Short: "controls the lnkshrtn server",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
		},
		DisableAutoGenTag: true,
	}

	command.AddCommand(commands.NewLinkCommand(&clientOpts))
	command.AddCommand(commands.NewContextCommand(&clientOpts))

	command.PersistentFlags().StringVar(&clientOpts.ServerAddr, "server", "", "lnkshrtn server address")
	command.PersistentFlags().StringVar(&clientOpts.ConfigPath, "config", "", "config file path")
	return command
}

//func overrideOptions() *apiclient.ApiClientOpts{
//	var ret apiclient.ApiClientOpts
//
//
//}
