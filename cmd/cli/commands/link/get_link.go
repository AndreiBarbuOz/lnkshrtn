package link

import (
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/spf13/cobra"
	"os"
)

func NewGetLinkCommand(clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	var command = &cobra.Command{
		Use:   "get",
		Short: "Get links",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
			os.Exit(1)
		},
	}
	return command
}
