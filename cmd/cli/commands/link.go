package commands

import (
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/spf13/cobra"
	"os"
)

func NewLinkCommand(clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	var command = &cobra.Command{
		Use:   "link",
		Short: "Manage links",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
			os.Exit(1)
		},
	}
	command.AddCommand(NewLinkListCommand(clientOpts))
	return command
}

func NewLinkListCommand(clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	var command = &cobra.Command{
		Use:   "list",
		Short: "List links",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
			os.Exit(1)
		},
	}
	return command
}
