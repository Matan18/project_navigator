package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"project_navigator.matandriola.com/database_data"
)

func Alias() *cobra.Command {
	return &cobra.Command{
		Use:     "alias [project_name] [alias]",
		Short:   "Add alias argument as a reference to project_name",
		Aliases: []string{"a", "add", "set", "s"},
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			database_data.RegisterAlias(args[1], args[0])

			fmt.Println("Alias registered")
		},
	}
}
