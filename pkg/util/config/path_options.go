package config

import (
	"os"
	"os/user"
	"path/filepath"
)

type PathOptions struct {
	// globalFile is the default file storing lnkshrtn configurations
	GlobalFile       string
	EnvVar           string
	ExplicitFileFlag string
}

type PathOptionsGetter interface {
	GetActualConfigFile() string
	IsExplicitFile() bool
	GetExplicitFile() string
	SetExplicitFlag(pathFlag string)
}

var _ PathOptionsGetter = &PathOptions{}

const (
	RecommendedConfigPathEnvVar = "LNKSHRTN_CONFIG"
	RecommendedHomeDir          = ".lnkshrtn"
	RecommendedFileName         = "config"
)

var (
	RecommendedConfigDir = filepath.Join(getHomeDir(), RecommendedHomeDir)
	RecommendedHomeFile  = filepath.Join(RecommendedConfigDir, RecommendedFileName)
)

func getHomeDir() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		usr, err := user.Current()
		if err != nil {
			return ""
		}
		homeDir = usr.HomeDir
	}
	return homeDir
}

func (p *PathOptions) getEnvVarFile() string {
	if len(p.EnvVar) == 0 {
		return ""
	}

	envVarValue := os.Getenv(p.EnvVar)
	if len(envVarValue) == 0 {
		return ""
	}

	return envVarValue
}

func (p *PathOptions) IsExplicitFile() bool {
	return len(p.ExplicitFileFlag) > 0
}

func (p *PathOptions) GetExplicitFile() string {
	return p.ExplicitFileFlag
}

func (p *PathOptions) GetActualConfigFile() string {
	if p.IsExplicitFile() {
		return p.GetExplicitFile()
	}
	envVarFile := p.getEnvVarFile()
	if len(envVarFile) > 0 {
		return envVarFile
	}
	return p.GlobalFile
}

func (p *PathOptions) SetExplicitFlag(pathFlag string) {
	p.ExplicitFileFlag = pathFlag
}

func NewDefaultPathOptions() *PathOptions {
	return &PathOptions{
		GlobalFile:       RecommendedHomeFile,
		EnvVar:           RecommendedConfigPathEnvVar,
		ExplicitFileFlag: "",
	}
}
