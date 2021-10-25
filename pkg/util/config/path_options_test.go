package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGetActualConfig_EnvVariable(t *testing.T) {
	if err := os.Setenv(RecommendedConfigPathEnvVar, "env-variable-file-path"); err != nil {
		t.Fatalf("Could not set env variable %s: %v", RecommendedConfigPathEnvVar, err)
	}
	defer os.Unsetenv(RecommendedConfigPathEnvVar)
	pathOptions := NewDefaultPathOptions()
	actual := pathOptions.GetActualConfigFile()
	assert.Equal(t, "env-variable-file-path", actual, "load order not observed: environment variable ignored")
}

func TestGetActualConfig_CliArgument(t *testing.T) {
	pathOptions := NewDefaultPathOptions()
	pathOptions.SetExplicitFlag("command-line-arg-file-path")
	actual := pathOptions.GetActualConfigFile()
	assert.Equal(t, "command-line-arg-file-path", actual, "load order not observed: cli argument ignored")
}

func TestGetActualConfig_HomeFolder(t *testing.T) {
	pathOptions := NewDefaultPathOptions()
	actual := pathOptions.GetActualConfigFile()
	assert.Contains(t, actual, filepath.Join(".lnkshrtn", "config"), "default config path not observed")
}

func TestGetActualConfig_CliArgumentEnv(t *testing.T) {
	if err := os.Setenv(RecommendedConfigPathEnvVar, "env-variable-file-path"); err != nil {
		t.Fatalf("Could not set env variable %s: %v", RecommendedConfigPathEnvVar, err)
	}
	defer os.Unsetenv(RecommendedConfigPathEnvVar)
	pathOptions := NewDefaultPathOptions()
	pathOptions.SetExplicitFlag("command-line-arg-file-path")
	actual := pathOptions.GetActualConfigFile()
	assert.Equal(t, "command-line-arg-file-path", actual, "load order not observed: cli argument ignored")
}
