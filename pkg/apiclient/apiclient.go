package apiclient

import "github.com/AndreiBarbuOz/lnkshrtn/pkg/util/config"

type ApiClientOpts struct {
	ServerAddr string
	ConfigPath config.PathOptions
	Context    string
	Headers    []string
}
