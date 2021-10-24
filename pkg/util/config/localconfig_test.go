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
	assert.NotNil(t, err, "decode failed to return error on illegal config")
}

func TestReadLocalConfig_NonExistingFile(t *testing.T) {
	var wd string
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get the working directory: %v", err)
	}
	config, err := ReadLocalConfig(filepath.Join(wd, "non-exiting-file"))
	assert.Nil(t, err, "decode failed to return error on illegal config")
	assert.NotNil(t, config, "Expected non nil config, but received nil")
}
