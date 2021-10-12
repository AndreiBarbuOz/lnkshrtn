package commands

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

func NewContextCommand(clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	var command = &cobra.Command{
		Use:   "context",
		Short: "Manage contexts",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
			os.Exit(1)
		},
	}
	command.AddCommand(NewContextListCommand(clientOpts))
	return command
}

func NewContextListCommand(clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	var localConfig *config.LocalConfig

	var command = &cobra.Command{
		Use:   "list",
		Short: "List contexts",
		Run: func(c *cobra.Command, args []string) {
			localConfigPath, err := config.DefaultLocalConfigPath()
			fmt.Printf(clientOpts.ServerAddr)
			if err != nil {
				log.Fatal(err)
			}
			localConfig, err = config.ReadLocalConfig(localConfigPath)
			if err != nil {
				log.Fatal(err)
			}
			if localConfig != nil {
				listContexts(localConfig)
			}
		},
	}
	return command
}

func listContexts(cfg *config.LocalConfig) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer func() { _ = w.Flush() }()
	columnNames := []string{"CURRENT", "NAME", "SERVER"}
	_, err := fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t"))

	if err != nil {
		log.Fatal(err)
	}

	for _, contextRef := range cfg.Contexts {
		context, err := cfg.ResolveContext(contextRef.Name)
		if err != nil {
			log.Fatalf("Context '%s' had error: %v", contextRef.Name, err)
		}
		prefix := " "
		if cfg.CurrentContext == context.Name {
			prefix = "*"
		}
		_, err = fmt.Fprintf(w, "%s\t%s\t%s\n", prefix, context.Name, context.Server.Server)
		if err != nil {
			log.Fatal(err)
		}
	}
}
