package main

import (
	"github.com/spf13/cobra"

	"project_navigator.matandriola.com/command"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(command.Alias())
	rootCmd.AddCommand(command.Register())
	rootCmd.AddCommand(command.FindDirectory())
	rootCmd.Execute()
}
