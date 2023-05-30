package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "uninstalledge",
		Short:   "Uninstall edge cli",
		Long:    "Uninstall mircosoft edge browser",
		Version: "0.0.1",
	}
)

func main() {
	rootCmd.Execute()
}
