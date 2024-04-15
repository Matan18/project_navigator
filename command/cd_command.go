package command

import (
	"os"

	"github.com/spf13/cobra"

	"project_navigator.matandriola.com/database_data"
)

func getPath(path string) string {
	alias_list := database_data.CheckAlias()

	path_value, ok := alias_list[path]
	if ok {
		return path_value
	}

	return path
}

func FindDirectory() *cobra.Command {
	return &cobra.Command{
		Use:     "find [project_name]",
		Aliases: []string{"f", "cd", "navigate", "nav", "go", "open", "o", "g"},
		Short:   "Navigate user to folder with project_name in registered folders",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path_list := database_data.CheckPath()

			path_arg := getPath(args[0])

			for _, path := range path_list {
				dirList, err := os.ReadDir(path)

				if err == nil {
					for _, dir := range dirList {
						if dir.IsDir() && dir.Name() == path_arg {
							os.Stdout.Write([]byte(path + path_arg))
						}
					}
				}
			}
		},
	}
}
