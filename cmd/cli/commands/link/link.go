package link

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
	command.AddCommand(NewGetLinkCommand(clientOpts))
	return command
}
