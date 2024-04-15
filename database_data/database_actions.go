package database_data

import (
	"os"
	"strings"
)

const FILE_PATH = "./folder_path_list"
const ALIAS_FILE_PATH = "./alias_list"

func readOrCreateFile() {
	_, err := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		_, err := os.Create(FILE_PATH)

		check(err)
	}
}

func readOrCreateAliasFile() {
	_, err := os.OpenFile(ALIAS_FILE_PATH, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		_, err := os.Create(ALIAS_FILE_PATH)

		check(err)
	}
}

func RegisterAlias(alias string, project string) {
	readOrCreateAliasFile()

	buffer, err := os.ReadFile(ALIAS_FILE_PATH)

	check(err)

	file_data := string(buffer)
	file_list := strings.Split(file_data, "\n")
	alias_exists := false

	for i := range file_list {
		if strings.Contains(file_list[i], alias) {
			alias_exists = true
			break
		}
	}

	if !alias_exists {
		new_file_list := append(file_list, alias+" "+project)
		file_string := strings.Join(new_file_list, "\n")

		err := os.WriteFile(ALIAS_FILE_PATH, []byte(file_string), 0644)

		check(err)
	}
}

func CheckAlias() map[string]string {
	readOrCreateAliasFile()

	buffer, err := os.ReadFile(ALIAS_FILE_PATH)

	check(err)

	file_data := string(buffer)
	file_list := strings.Split(file_data, "\n")
	alias_map := make(map[string]string)

	for i := range file_list {
		alias := strings.Split(file_list[i], " ")

		if len(alias) == 2 {
			alias_map[alias[0]] = alias[1]
		}
	}

	return alias_map
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckPath() []string {
	readOrCreateFile()

	buffer, err := os.ReadFile(FILE_PATH)

	check(err)

	file_data := string(buffer)
	file_list := strings.Split(file_data, "\n")
	path_list := make([]string, len(file_list))

	for i := range file_list {
		exists := false

		for j := range file_list {
			if file_list[i] == path_list[j] {
				exists = true
				break
			}
		}

		if !exists {
			path_list = append(path_list, file_list[i])
		}
	}

	return path_list
}

func save_path(path_list []string) {
	file_string := strings.Join(path_list, "\n")

	err := os.WriteFile(FILE_PATH, []byte(file_string), 0644)

	check(err)
}

func RegisterPath(path string) {
	file_list := CheckPath()
	new_file_list := append(file_list, path)

	save_path(new_file_list)
}
