# Project navigator

My intention with this project is to turn easy to navigate through the projects of my work using CLI, with some easy actions like:

- `navigate register ~/project/node/react`: this command will register the path to something like `source folders`;
- `navigate code kanban`: this command will open vscode in the project with folder name like `kanban` inside my `source folders`;
- `navigate cd kanban`: this command will just navigate to the project folder named like `kanban`;
- `navigate alias kanban kb`: this command will register `kb` as an alias for the `kanban` project;
- `navigate cd kb`: this command will open vscode in the project with folder name like `kanban` (since has already registered as an alias);

## Tech

- [Go Lang](https://go.dev/)
- [cobra](https://cobra.dev/)
