package config

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/kballard/go-shellquote"
)

type EnvFlags struct {
	flags map[string]string
}

func NewEnvFlags() *EnvFlags {
	var ret EnvFlags
	err := ret.loadFlags()
	if err != nil {
		log.Fatal(err)
	}
	return &ret
}

func (e *EnvFlags) loadFlags() error {
	e.flags = make(map[string]string)

	opts, err := shellquote.Split(os.Getenv("LNKSHRTN_OPTS"))
	if err != nil {
		return err
	}

	// crtFlag allows for key-value arguments
	var crtFlag string
	for _, opt := range opts {
		if strings.HasPrefix(opt, "--") {
			if crtFlag != "" {
				e.flags[crtFlag] = "true"
			}
			crtFlag = strings.TrimPrefix(opt, "--")
		} else if crtFlag != "" {
			e.flags[crtFlag] = opt
			crtFlag = ""
		} else {
			return errors.New("LNKSHRTN_OPTS invalid at '" + opt + "'")
		}
	}
	if crtFlag != "" {
		e.flags[crtFlag] = "true"
	}
	return nil
}

func (e *EnvFlags) GetFlag(key, fallback string) string {
	val, ok := e.flags[key]
	if ok {
		return val
	}
	return fallback
}

func (e *EnvFlags) GetBoolFlag(key string) bool {
	return e.GetFlag(key, "false") == "true"
}
