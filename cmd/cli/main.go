package main

import (
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/cmd/cli/commands/config"
	"github.com/AndreiBarbuOz/lnkshrtn/cmd/cli/commands/link"
	"github.com/AndreiBarbuOz/lnkshrtn/cmd/cli/util"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/spf13/cobra"
	"os"
)

const (
	cliName = "lnkctl"
)

func main() {
	var command *cobra.Command

	ioStreams := util.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	command = NewCommand(ioStreams)

	err := command.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewCommand(ioStreams util.IOStreams) *cobra.Command {

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

	command.AddCommand(link.NewLinkCommand(ioStreams, &clientOpts))
	command.AddCommand(config.NewConfigCommand(ioStreams, &clientOpts))

	command.PersistentFlags().StringVar(&clientOpts.ConfigPath.ExplicitFileFlag, "config", "", "config file path")
	return command
}

//func overrideOptions() *apiclient.ApiClientOpts{
//	var ret apiclient.ApiClientOpts
//
//
//}
