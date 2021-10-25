package config

import (
	"os"
	"os/user"
	"path/filepath"
)

type PathOptions struct {
	globalFile       string
	envVar           string
	explicitFileFlag string
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
	if len(p.envVar) == 0 {
		return ""
	}

	envVarValue := os.Getenv(p.envVar)
	if len(envVarValue) == 0 {
		return ""
	}

	return envVarValue
}

func (p *PathOptions) IsExplicitFile() bool {
	return len(p.explicitFileFlag) > 0
}

func (p *PathOptions) GetExplicitFile() string {
	return p.explicitFileFlag
}

func (p *PathOptions) GetActualConfigFile() string {
	if p.IsExplicitFile() {
		return p.GetExplicitFile()
	}
	envVarFile := p.getEnvVarFile()
	if len(envVarFile) > 0 {
		return envVarFile
	}
	return p.globalFile
}

func (p *PathOptions) SetExplicitFlag(pathFlag string) {
	p.explicitFileFlag = pathFlag
}

func NewDefaultPathOptions() *PathOptions {
	return &PathOptions{
		globalFile:       RecommendedHomeFile,
		envVar:           RecommendedConfigPathEnvVar,
		explicitFileFlag: "",
	}
}
