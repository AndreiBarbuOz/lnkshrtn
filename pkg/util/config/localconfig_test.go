package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestLocalConfig_ReadLocalConfig(t *testing.T) {
	filePath := filepath.Join("../../../test/testdata", "config.yaml")
	absolutePath, _ := filepath.Abs(filePath)
	fmt.Printf("%+v\n", absolutePath)
	cfg, err := ReadLocalConfig(absolutePath)
	assert.Nil(t, err, "ReadLocalConfig returned non nil error")
	assert.NotNil(t, cfg, "ReadLocalConfig returned nil config for existing file")

	assert.Equal(t, cfg.CurrentContext, "dev-test.example.com",
		"ReadLocalConfig returned wrong current context")
	assert.Contains(t, cfg.Servers, Server{Server: "dev-test.example.com:8080"})
	assert.Contains(t, cfg.Contexts, ContextRef{
		Name:   "dev-test.example.com",
		Server: "dev-test.example.com:8080",
	})
}

func TestLocalConfig_ReadLocalConfig_nonexisting(t *testing.T) {
	filePath := filepath.Join("../../../test/testdata", "config-foo.yaml")
	absolutePath, _ := filepath.Abs(filePath)
	fmt.Printf("%+v\n", absolutePath)
	cfg, err := ReadLocalConfig(absolutePath)
	assert.Nil(t, err, "ReadLocalConfig returned non nil error for non existing file")
	assert.Nil(t, cfg, "ReadLocalConfig returned non nil cfg for non existing file")
}

func TestLocalConfig_ReadLocalConfig_invalid(t *testing.T) {
	filePath := filepath.Join("../../../test/testdata", "config-invalid.yaml")
	absolutePath, _ := filepath.Abs(filePath)
	fmt.Printf("%+v\n", absolutePath)
	cfg, err := ReadLocalConfig(absolutePath)
	assert.NotNil(t, err, "ReadLocalConfig returned nil error for invalid file")
	assert.Nil(t, cfg, "ReadLocalConfig returned non nil cfg for invalid file")
}
