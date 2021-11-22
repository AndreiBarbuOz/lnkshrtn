package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDecode_NilErr(t *testing.T) {
	var testData = `contexts:
  - name: alpha
    server: server-alpha
current-context: alpha
servers:
  - server: server-alpha
`

	_, err := decode([]byte(testData))
	assert.Nil(t, err, "decode returned non nil error: %v", err)
}

func TestDecode_Invalid(t *testing.T) {
	var testData = `key: value`
	_, err := decode([]byte(testData))
	assert.NotNil(t, err, "decode failed to return error on illegal config")
}

func TestDecode_ExtraFields(t *testing.T) {
	var testData = `contexts:
  - name: alpha
    server: server-alpha
	key: value
current-context: alpha
servers:
  - server: server-alpha`
	_, err := decode([]byte(testData))
	assert.NotNil(t, err, "decode failed to return error on illegal config")
}

func TestReadLocalConfig_NilErr(t *testing.T) {
	cliFile, _ := ioutil.TempFile("", "")
	defer os.Remove(cliFile.Name())

	if err := ioutil.WriteFile(cliFile.Name(), []byte("illegal value"), 0644); err != nil {
		t.Fatalf("Error creating tempfile: %v", err)
	}
	_, err := ReadLocalConfig(cliFile.Name())
	assert.NotNil(t, err, "ReadLocalConfig failed to return error on illegal config")
}

func TestReadLocalConfig_NonExistingFile(t *testing.T) {
	var wd string
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get the working directory: %v", err)
	}
	config, err := ReadLocalConfig(filepath.Join(wd, "non-exiting-file"))
	assert.Nil(t, err, "decode failed to return nil error on non existing config")
	assert.NotNil(t, config, "Expected non nil config, but received nil")
}

func TestReadLocalConfig_ValidFile(t *testing.T) {
	cliFile, _ := ioutil.TempFile("", "")
	defer os.Remove(cliFile.Name())

	if err := ioutil.WriteFile(cliFile.Name(), []byte(`contexts:
  - name: alpha
    server: server-alpha
current-context: alpha
servers:
  - server: server-alpha`), 0644); err != nil {
		t.Fatalf("Error creating tempfile: %v", err)
	}
	actual, err := ReadLocalConfig(cliFile.Name())
	assert.Nil(t, err, "decode failed to return nil error on valid config")

	var expected = &LocalConfig{
		CurrentContext: "alpha",
		Contexts: []Context{
			{
				Name:   "alpha",
				Server: "server-alpha",
			},
		},
		Servers: []Server{
			{
				Server: "server-alpha",
			},
		},
	}
	assert.Equal(t, expected, actual, "config read from file does not match expected")
}

func TestReadLocalConfig_ValidFileMultipleContexts(t *testing.T) {
	cliFile, _ := ioutil.TempFile("", "")
	defer os.Remove(cliFile.Name())

	if err := ioutil.WriteFile(cliFile.Name(), []byte(`contexts:
  - name: alpha
    server: server-alpha
  - name: beta
    server: server-beta
current-context: beta
servers:
  - server: server-alpha
  - server: server-beta`), 0644); err != nil {
		t.Fatalf("Error creating tempfile: %v", err)
	}
	actual, err := ReadLocalConfig(cliFile.Name())
	assert.Nil(t, err, "decode failed to return nil error on valid config")

	var expected = &LocalConfig{
		CurrentContext: "beta",
		Contexts: []Context{
			{
				Name:   "alpha",
				Server: "server-alpha",
			},
			{
				Name:   "beta",
				Server: "server-beta",
			},
		},
		Servers: []Server{
			{
				Server: "server-alpha",
			},
			{
				Server: "server-beta",
			},
		},
	}
	assert.Equal(t, expected, actual, "config read from file does not match expected")
}

func TestReadLocalConfig_InvalidContext(t *testing.T) {
	cliFile, _ := ioutil.TempFile("", "")
	defer os.Remove(cliFile.Name())

	if err := ioutil.WriteFile(cliFile.Name(), []byte(`contexts:
  - name: alpha
    server: server-alpha
current-context: beta
servers:
  - server: server-alpha`), 0644); err != nil {
		t.Fatalf("Error creating tempfile: %v", err)
	}
	actual, err := ReadLocalConfig(cliFile.Name())
	assert.Nil(t, actual, "ReadLocalConfig failed to return nil output for invalid context")
	assert.NotNil(t, err, "ReadLocalConfig failed to return non nil error for invalid context")
}

func TestReadLocalConfig_AfterWrite(t *testing.T) {
	cliFile, _ := ioutil.TempFile("", "")
	defer os.Remove(cliFile.Name())

	var cfg = &LocalConfig{
		CurrentContext: "beta",
		Contexts: []Context{
			{
				Name:   "alpha",
				Server: "server-alpha",
			},
			{
				Name:   "beta",
				Server: "server-beta",
			},
		},
		Servers: []Server{
			{
				Server: "server-alpha",
			},
			{
				Server: "server-beta",
			},
		},
	}

	err := WriteConfigToFile(cfg, cliFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual, err := ReadLocalConfig(cliFile.Name())
	assert.Nil(t, err, "decode failed to return nil error on valid config")

	assert.Equal(t, cfg, actual, "config read from file does not match expected")
}
