package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

// LocalConfig is a local lnkshrtn config file
type LocalConfig struct {
	CurrentContext string       `yaml:"current-context"`
	Contexts       []ContextRef `yaml:"contexts"`
	Servers        []Server     `yaml:"servers"`
}

// ContextRef is a reference to a Server and User for an API client
type ContextRef struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

// Context is the resolved Server object
type Context struct {
	Name   string
	Server Server
}

// Server contains the lnkshrtn server information
type Server struct {
	// Server is the lnkshrtn server address
	Server string `yaml:"server"`
}

// ReadLocalConfig loads up the local configuration file. Returns nil if config does not exist
func ReadLocalConfig(path string) (*LocalConfig, error) {
	var err error
	var config LocalConfig

	err = UnmarshalLocalConfigFile(path, &config)
	if os.IsNotExist(err) {
		return nil, nil
	}
	err = ValidateLocalConfig(config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// ValidateLocalConfig validates the configuration object, after it was unmarshalled
func ValidateLocalConfig(config LocalConfig) error {
	if config.CurrentContext == "" {
		return nil
	}
	if _, err := config.ResolveContext(config.CurrentContext); err != nil {
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

// ResolveContext resolves the specified context. If unspecified, resolves the current context
func (l *LocalConfig) ResolveContext(name string) (*Context, error) {
	if name == "" {
		if l.CurrentContext == "" {
			return nil, fmt.Errorf("local config: current-context unset")
		}
		name = l.CurrentContext
	}
	for _, ctx := range l.Contexts {
		if ctx.Name == name {
			server, err := l.GetServer(ctx.Server)
			if err != nil {
				return nil, err
			}
			return &Context{
				Name:   ctx.Name,
				Server: *server,
			}, nil
		}
	}
	return nil, fmt.Errorf("context '%s' undefined", name)
}

func (l *LocalConfig) GetServer(name string) (*Server, error) {
	for _, s := range l.Servers {
		if s.Server == name {
			return &s, nil
		}
	}
	return nil, fmt.Errorf("server '%s' undefined", name)
}

func (l *LocalConfig) UpsertServer(server Server) {
	for i, s := range l.Servers {
		if s.Server == server.Server {
			l.Servers[i] = server
			return
		}
	}
	l.Servers = append(l.Servers, server)
}

// Returns true if server was removed successfully
func (l *LocalConfig) RemoveServer(serverName string) bool {
	for i, s := range l.Servers {
		if s.Server == serverName {
			l.Servers = append(l.Servers[:i], l.Servers[i+1:]...)
			return true
		}
	}
	return false
}

func (l *LocalConfig) UpsertContext(context ContextRef) {
	for i, c := range l.Contexts {
		if c.Name == context.Name {
			l.Contexts[i] = context
			return
		}
	}
	l.Contexts = append(l.Contexts, context)
}

// Returns true if context was removed successfully
func (l *LocalConfig) RemoveContext(serverName string) (string, bool) {
	for i, c := range l.Contexts {
		if c.Name == serverName {
			l.Contexts = append(l.Contexts[:i], l.Contexts[i+1:]...)
			return c.Server, true
		}
	}
	return "", false
}

func (l *LocalConfig) IsEmpty() bool {
	return len(l.Servers) == 0
}

// DefaultConfigDir returns the local configuration path
func DefaultConfigDir() (string, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		homeDir = usr.HomeDir
	}
	return path.Join(homeDir, ".lnkshrtn"), nil
}

// DefaultLocalConfigPath returns the default local configuration path
func DefaultLocalConfigPath() (string, error) {
	dir, err := DefaultConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(dir, "config"), nil
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

// UnmarshalLocalConfigFile retrieves JSON or YAML from a file on disk.
// The caller is responsible for checking error return values.
func UnmarshalLocalConfigFile(path string, obj interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		return fmt.Errorf("failed to read file %s", path)
	}
	err = yaml.Unmarshal(data, obj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal object: %w", err)
	}

	return nil
}
