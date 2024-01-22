package main

import (
	"fmt"
	"github.com/rsturla/platform/tools/platform-cli/internal/command"
	"os"
)

func main() {
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
