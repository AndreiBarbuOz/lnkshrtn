package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"

	server "github.com/AndreiBarbuOz/lnkshrtn/cmd/server/commands"
)

const (
	binaryNameEnv = "LNKSHRTN_BINARY_NAME"
)

func main() {
	var command *cobra.Command

	binaryName := filepath.Base(os.Args[0])
	if val := os.Getenv(binaryNameEnv); val != "" {
		binaryName = val
	}
	switch binaryName {
	case "lnk-server":
		command = server.NewCommand()
	}
	if command == nil {
		fmt.Printf("Unknown binary name %s\n", binaryName)
		os.Exit(1)
	}
	err := command.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
