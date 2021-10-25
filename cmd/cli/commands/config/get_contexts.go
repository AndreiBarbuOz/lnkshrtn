package config

import (
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/util/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func NewGetContextsCommand(clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	var pathOptions *config.PathOptions
	pathOptions = config.NewDefaultPathOptions()
	if len(clientOpts.Context) == 0 {
		pathOptions.SetExplicitFlag(clientOpts.Context)
	}

	var command = &cobra.Command{
		Use:   "get",
		Short: "Get contexts",
		Run: func(c *cobra.Command, args []string) {
			configPath := pathOptions.GetActualConfigFile()
			fmt.Printf(clientOpts.ServerAddr)
			resolvedConfig, err := config.ReadLocalConfig(configPath)
			if err != nil {
				log.Fatal(err)
			}
			if resolvedConfig != nil {
				getContexts(resolvedConfig)
			}
		},
	}
	return command
}

func getContexts(cfg *config.LocalConfig) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer func() { _ = w.Flush() }()
	columnNames := []string{"CURRENT", "NAME", "SERVER"}
	_, err := fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t"))

	if err != nil {
		log.Fatal(err)
	}

	for _, crt := range cfg.Contexts {

		prefix := " "
		if cfg.CurrentContext == crt.Name {
			prefix = "*"
		}
		_, err = fmt.Fprintf(w, "%s\t%s\t%s\n", prefix, crt.Name, crt.Server)
		if err != nil {
			log.Fatal(err)
		}
	}
}
