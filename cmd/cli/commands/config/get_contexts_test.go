package config

import (
	"github.com/AndreiBarbuOz/lnkshrtn/cmd/cli/util"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/apiclient"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/util/config"
	"io/ioutil"
	"os"
	"testing"
)

type getContextsTest struct {
	cfg            config.LocalConfig
	expectedOutput string
	noHeaders      bool
	nameOnly       bool
	names          []string
}

func TestGetContextsAll(t *testing.T) {
	testConf := config.LocalConfig{
		CurrentContext: "dragon",
		Contexts: []config.Context{
			{
				Name:   "dragon",
				Server: "slayer",
			},
		},
		Servers: []config.Server{
			{
				Server: "slayer",
			},
		},
	}
	test := getContextsTest{
		cfg:       testConf,
		noHeaders: false,
		nameOnly:  false,
		names:     []string{},
		expectedOutput: `CURRENT  NAME    SERVER
*        dragon  slayer
`,
	}
	test.run(t)
}

func TestGetContextsNoHeaders(t *testing.T) {
	testConf := config.LocalConfig{
		CurrentContext: "dragon",
		Contexts: []config.Context{
			{
				Name:   "dragon",
				Server: "slayer",
			},
		},
		Servers: []config.Server{
			{
				Server: "slayer",
			},
		},
	}
	test := getContextsTest{
		cfg:       testConf,
		noHeaders: true,
		nameOnly:  false,
		names:     []string{},
		expectedOutput: `*  dragon  slayer
`,
	}
	test.run(t)
}

func TestGetContextsOnlyNames(t *testing.T) {
	testConf := config.LocalConfig{
		CurrentContext: "dragon",
		Contexts: []config.Context{
			{
				Name:   "dragon",
				Server: "slayer",
			},
		},
		Servers: []config.Server{
			{
				Server: "slayer",
			},
		},
	}
	test := getContextsTest{
		cfg:       testConf,
		noHeaders: false,
		nameOnly:  true,
		names:     []string{},
		expectedOutput: `dragon
`,
	}
	test.run(t)
}

func TestGetContextsEmpty(t *testing.T) {
	testConf := config.LocalConfig{}
	test := getContextsTest{
		cfg:            testConf,
		noHeaders:      true,
		nameOnly:       false,
		names:          []string{},
		expectedOutput: "",
	}
	test.run(t)
}

func TestGetContextsMultiple(t *testing.T) {
	testConf := config.LocalConfig{
		CurrentContext: "dragon",
		Contexts: []config.Context{
			{
				Name:   "dragon",
				Server: "slayer",
			},
			{
				Name:   "saint",
				Server: "george",
			},
		},
		Servers: []config.Server{
			{
				Server: "slayer",
			},
			{
				Server: "george",
			},
		},
	}
	test := getContextsTest{
		cfg:       testConf,
		noHeaders: false,
		nameOnly:  true,
		names:     []string{},
		expectedOutput: `dragon
saint
`,
	}
	test.run(t)
}

func TestGetContextsMultipleArgs(t *testing.T) {
	testConf := config.LocalConfig{
		CurrentContext: "dragon",
		Contexts: []config.Context{
			{
				Name:   "dragon",
				Server: "slayer",
			},
			{
				Name:   "saint",
				Server: "george",
			},
		},
		Servers: []config.Server{
			{
				Server: "slayer",
			},
			{
				Server: "george",
			},
		},
	}
	test := getContextsTest{
		cfg:       testConf,
		noHeaders: false,
		nameOnly:  false,
		names:     []string{"dragon", "saint"},
		expectedOutput: `CURRENT  NAME    SERVER
*        dragon  slayer
         saint   george
`,
	}
	test.run(t)
}

func (test getContextsTest) run(t *testing.T) {
	configFile, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(configFile.Name())
	err = config.WriteConfigToFile(&test.cfg, configFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pathOptions := config.NewDefaultPathOptions()
	pathOptions.SetExplicitFlag(configFile.Name())
	ioStreams, _, buf, _ := util.NewTestIOStreams()
	options := apiclient.ApiClientOpts{
		ServerAddr: "",
		ConfigPath: *pathOptions,
		Context:    "",
		Headers:    nil,
	}

	cmd := NewGetContextsCommand(ioStreams, &options)
	if test.nameOnly {
		cmd.Flags().Set("output", "name")
	}
	if test.noHeaders {
		cmd.Flags().Set("no-headers", "true")
	}
	cmd.Run(cmd, test.names)
	if len(test.expectedOutput) != 0 {
		if buf.String() != test.expectedOutput {
			t.Errorf("Expected\n%s\ngot\n%s", test.expectedOutput, buf.String())
		}
		return
	}
}
