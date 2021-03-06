package config

import (
	"github.com/AndreiBarbuOz/lnkshrtn/cmd/cli/util"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/spf13/cobra"
	"os"
)

func NewConfigCommand(ioStreams util.IOStreams, clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	var command = &cobra.Command{
		Use:   "config",
		Short: "Manage configurations",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
			os.Exit(1)
		},
	}
	command.AddCommand(NewGetContextsCommand(ioStreams, clientOpts))
	return command
}
