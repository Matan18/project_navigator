package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"project_navigator.matandriola.com/database_data"
)

func Register() *cobra.Command {
	return &cobra.Command{
		Use:   "register [path]",
		Short: "Add path to a list of folders to search projects",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			database_data.RegisterPath(path)
			fmt.Printf("Path %s registered\n", path)
		},
	}
}
