package config

import (
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/cmd/cli/util"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/util/config"
	"github.com/spf13/cobra"
	"io"
	"log"
	"sort"
	"strings"
	"text/tabwriter"
)

type GetContextsOptions struct {
	clientOpts   apiclient.ApiClientOpts
	nameOnly     bool
	showHeaders  bool
	contextNames []string

	ioStreams util.IOStreams
}

func NewGetContextsCommand(ioStreams util.IOStreams, clientOpts *apiclient.ApiClientOpts) *cobra.Command {
	options := &GetContextsOptions{
		clientOpts: *clientOpts,
		ioStreams:  ioStreams,
	}
	var pathOptions *config.PathOptions
	pathOptions = config.NewDefaultPathOptions()
	if len(clientOpts.Context) == 0 {
		pathOptions.SetExplicitFlag(clientOpts.Context)
	}

	var command = &cobra.Command{
		Use:   "get",
		Short: "Get contexts",
		Run: func(c *cobra.Command, args []string) {
			options.contextNames = args
			output, _ := c.Flags().GetString("output")
			if output == "name" {
				options.nameOnly = true
			}
			options.showHeaders = true
			noHeader, _ := c.Flags().GetBool("no-header")
			if noHeader || options.nameOnly {
				options.showHeaders = false
			}

			err := options.RunGetContexts()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	command.Flags().Bool("no-headers", false, "When using the default output format, don't print headers.")
	command.Flags().StringP("output", "o", "", "Output format. One of: name")
	return command
}

// RunGetContexts implements all the necessary functionality for context retrieval.
func (o GetContextsOptions) RunGetContexts() error {
	configPath := o.clientOpts.ConfigPath.GetActualConfigFile()
	cfg, err := config.ReadLocalConfig(configPath)
	if err != nil {
		return fmt.Errorf("could not read local config at %s: %w", configPath, err)
	}

	out, found := o.ioStreams.Out.(*tabwriter.Writer)
	if !found {
		out = tabwriter.NewWriter(o.ioStreams.Out, 0, 0, 2, ' ', 0)
		defer func() { _ = out.Flush() }()
	}

	// Build a list of context names to print, and warn if any requested contexts are not found.
	// Do this before printing the headers so it doesn't look ugly.
	var allErrs []error
	var toPrint []string
	if len(o.contextNames) == 0 {
		for _, cxt := range cfg.Contexts {
			toPrint = append(toPrint, cxt.Name)
		}
	} else {
		for _, name := range o.contextNames {
			found = false
			for _, cxt := range cfg.Contexts {
				if name == cxt.Name {
					toPrint = append(toPrint, name)
					found = true
					break
				}
			}
			if !found {
				allErrs = append(allErrs, fmt.Errorf("context %v not found", name))
			}
		}
	}
	if o.showHeaders {
		err = printContextHeaders(out, o.nameOnly)
		if err != nil {
			allErrs = append(allErrs, err)
		}
	}

	sort.Strings(toPrint)
	for _, name := range toPrint {
		var crtCtx *config.Context
		for _, ctx := range cfg.Contexts {
			if name == ctx.Name {
				crtCtx = &ctx
				break
			}
		}

		err = printContext(name, crtCtx, out, o.nameOnly, cfg.CurrentContext == name)
		if err != nil {
			allErrs = append(allErrs, err)
		}
	}

	if len(allErrs) != 0 {
		return allErrs[0]
	}
	return nil
}

func printContextHeaders(out io.Writer, nameOnly bool) error {
	columnNames := []string{"CURRENT", "NAME", "SERVER"}
	if nameOnly {
		columnNames = columnNames[:1]
	}
	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columnNames, "\t"))
	return err
}

func printContext(name string, context *config.Context, w io.Writer, nameOnly, current bool) error {
	if nameOnly {
		_, err := fmt.Fprintf(w, "%s\n", name)
		return err
	}
	prefix := " "
	if current {
		prefix = "*"
	}
	_, err := fmt.Fprintf(w, "%s\t%s\t%s\n", prefix, name, context.Server)
	return err
}
