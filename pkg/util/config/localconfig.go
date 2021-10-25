package config

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// LocalConfig is a local lnkshrtn config file
type LocalConfig struct {
	CurrentContext string    `yaml:"current-context"`
	Contexts       []Context `yaml:"contexts"`
	Servers        []Server  `yaml:"servers"`
}

// Context is a reference to a lnkshrt Server
type Context struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

// Server contains the lnkshrtn server information
type Server struct {
	// Server is the lnkshrtn server address
	Server string `yaml:"server"`
}

func NewLocalConfig() *LocalConfig {
	return &LocalConfig{
		CurrentContext: "",
		Contexts:       make([]Context, 0, 0),
		Servers:        make([]Server, 0, 0),
	}
}

// ReadLocalConfig loads up the local configuration file. Returns nil if config does not exist
func ReadLocalConfig(path string) (*LocalConfig, error) {
	var err error
	var ret *LocalConfig

	ret, err = getConfigFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read local config %s: %w", path, err)
	}
	if ret == nil {
		return NewLocalConfig(), nil
	}
	err = validateLocalConfig(*ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// validateLocalConfig validates the configuration object, after it was unmarshalled
func validateLocalConfig(config LocalConfig) error {
	if config.CurrentContext == "" {
		return nil
	}
	if _, err := config.resolveContext(config.CurrentContext); err != nil {
		return fmt.Errorf("local config invalid: %w", err)
	}
	return nil
}

// WriteLocalConfig writes a new local configuration file.
func WriteLocalConfig(config LocalConfig, configPath string) error {
	err := os.MkdirAll(path.Dir(configPath), os.ModePerm)
	if err != nil {
		return err
	}
	return MarshalLocalYAMLFile(configPath, config)
}

func DeleteLocalConfig(configPath string) error {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return err
	}
	return os.Remove(configPath)
}

// resolveContext resolves the specified context. If unspecified, resolves the current context
func (l *LocalConfig) resolveContext(name string) (*Context, error) {
	if name == "" {
		if l.CurrentContext == "" {
			return nil, fmt.Errorf("local config: current-context unset")
		}
		name = l.CurrentContext
	}
	for _, ctx := range l.Contexts {
		if ctx.Name == name {
			server, err := l.getServer(ctx.Server)
			if err != nil {
				return nil, err
			}
			return &Context{
				Name:   ctx.Name,
				Server: server.Server,
			}, nil
		}
	}
	return nil, fmt.Errorf("context '%s' undefined", name)
}

func (l *LocalConfig) getServer(name string) (*Server, error) {
	if name == "" {
		return nil, fmt.Errorf("server name cannot be empty")
	}
	for _, srv := range l.Servers {
		if srv.Server == name {
			return &srv, nil
		}
	}
	return nil, fmt.Errorf("server %s not found", name)
}

// MarshalLocalYAMLFile writes JSON or YAML to a file on disk.
// The caller is responsible for checking error return values.
func MarshalLocalYAMLFile(path string, obj interface{}) error {
	yamlData, err := yaml.Marshal(obj)
	if err == nil {
		err = ioutil.WriteFile(path, yamlData, 0600)
	}
	return err
}

// getConfigFromFile retrieves YAML from a file on disk.
// The caller is responsible for checking error return values.
func getConfigFromFile(path string) (*LocalConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return decode(data)
}

func decode(data []byte) (*LocalConfig, error) {
	var ret LocalConfig
	var reader io.Reader = bytes.NewReader(data)
	var decoder = yaml.NewDecoder(reader)
	decoder.KnownFields(true)

	err := decoder.Decode(&ret)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %w", err)
	}
	return &ret, nil
}
