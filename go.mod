module project_navigator.matandriola.com/project_navigator

go 1.20

require github.com/spf13/cobra v1.8.0

require project_navigator.matandriola.com/database_data v0.0.0-00010101000000-000000000000 // indirect

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	project_navigator.matandriola.com/command v0.0.0-00010101000000-000000000000
)

replace project_navigator.matandriola.com/command => ./command

replace project_navigator.matandriola.com/database_data => ./database_data
