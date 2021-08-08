package commands

import (
	"context"
	apiServer "github.com/AndreiBarbuOz/lnkshrtn/pkg/links"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var command = cobra.Command{
		Use:               "lnkshrtn-sidecar",
		Short:             "Run Lnkshrtn sidecar",
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			for {
				ctx := context.Background()
				lnkshrtn := apiServer.NewServer(ctx)
				lnkshrtn.Run(ctx, 8080)
			}
		},
	}
	return &command
}
