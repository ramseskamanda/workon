# workon

## Usage

```zsh
workon [--project|-p <project>] [--vscode] [--latest]

When searching for a project, if more than one is found with the same name, it should show a picker with full path to both.
```

## Base Features

- [x] Basic command line arguments parsing
- [x] Determine what environment we're in (iTerm2 or Terminal)
- [x] Search for a project by name
- [x] Open project in VSCode when found
- [x] Easy installation on any PC
- [ ] Open project in NeoVim when found

## Future features

- [x] Start a new project
- [ ] Sort the projects by most recently accessed
- [ ] Welcome message for new users
- [ ] Exclude certain directories from file crawling
- [ ] Allow for easy extensions (e.g. startup scripts to source a virtualenv when opening a python project)
- [ ] Command line suggestions when going through questions

## How it works

- [x] 1 - Load configuration (~/.config):
  - [x] 1.1 - Determine what environment we're running in
  - [x] 1.2 - If config exists, skip to 2
  - [x] 1.3 - If not, start initialisation
  - [x] 1.4 - Ask for Documents path
  - [x] 1.5 - Ask for preferred editor (can autodetect present editors)
  - [x] 1.6 - If neovim, vim, or terminal, ask what profile to launch with
  - [x] 1.7 - Write config file
  - [x] 1.7 - Check CLI arguments, if project is present, skip to 5 with default editor
- [x] 2 - Load projects db:
  - [x] 2.1 - Start scanning goroutine for Documents path
  - [x] 2.2 - Wait for goroutine to finish
  - [x] 2.3 - Read db for list of projects
- [x] 3 - Prompt the user to search for a project to work on
  - [x] 3.1 - If user search doesn't get anywhere, ask if they want to create the project
  - [x] 3.2 - If user search is valid, skip to 4
- [x] 4 - Prompt the user for an editor (choice widget with default from config)
- [ ] 5 - Launch the project
  - [x] 5.1 - Check the editor type
  - [ ] 5.2 - If neovim, vim, or terminal, launch with profile from config file and cwd
  - [x] 5.3 - If code, or others, launch with all defaults
