package system

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func AskInitializationQuestions(cfg *Config) error {
	directory_prompt := &promptui.Prompt{
		Label:     "Root Project Directory (full path)",
		Default:   "/",
		IsVimMode: true,
		AllowEdit: true,
	}
	dir, err := directory_prompt.Run()
	if err != nil {
		return err
	}
	editor_prompt := promptui.Select{
		Label:     "Preferred Editor",
		IsVimMode: true,
		Items:     []string{string(VSCODE), string(NEOVIM), string(VIM), string(EMACS)},
	}
	_, editor, err := editor_prompt.Run()
	if err != nil {
		return err
	}
	editor_profile_prompt := promptui.Select{
		Label:     "Preferred Profile (for Terminal-based environments)",
		IsVimMode: true,
		Items:     []string{string(VSCODE), string(NEOVIM), string(VIM), string(EMACS)},
	}
	_, editor_profile, err := editor_profile_prompt.Run()
	if err != nil {
		return err
	}

	cfg = &Config{
		Path:               dir,
		Editor:             editor,
		Profile:            editor_profile,
		SupportedLanguages: []string{"python", "go", "flutter"},
	}

	return nil
}

func AskForProject(index *InternalIndex) (string, error) {
	directory_prompt := &promptui.Prompt{
		Label:     fmt.Sprintf("What do you want to work on? (Enter '%s' to start a new project)", NEW_PROJECT_PROMPT),
		IsVimMode: true,
		Validate: func(s string) error {
			if s == "__new" {
				return nil
			}
			for _, recent := range index.MostRecent {
				if strings.EqualFold(s, recent) {
					return nil
				}
			}

			_, err := index.FindProjectByName(s)
			perr, notok := err.(*ProjectNotFound)
			if notok {
				return perr
			} else {
				return nil
			}
		},
	}
	return directory_prompt.Run()
}

func StartNewProject(cfg *Config, index *InternalIndex) (string, error) {
	name_prompt := &promptui.Prompt{
		Label:     "What do you want to name your new project?",
		IsVimMode: true,
		Default:   "UntitledProject",
		Validate: func(s string) error {
			_, err := index.FindProjectByName(s)
			_, ok := err.(*ProjectNotFound)
			if ok {
				return nil
			} else {
				return errors.New("Project already exists")
			}
		},
	}
	name, err := name_prompt.Run()
	if err != nil {
		return name, err
	}

	project_type_prompt := &promptui.Select{
		Label:     "Project type",
		IsVimMode: true,
		Items:     cfg.SupportedLanguages,
	}

	_, project_type, err := project_type_prompt.Run()
	if err != nil {
		return project_type, err
	}

	path_prompt := &promptui.Prompt{
		Label:     "Project Path",
		Default:   cfg.Path,
		AllowEdit: true,
		IsVimMode: true,
	}

	path, err := path_prompt.Run()
	if err != nil {
		return path, err
	}

	err = CreateProject(name, path, project_type)
	if err != nil {
		return "", err
	}

	return name, nil
}

func AskForEditor() (string, error) {
	editor_prompt := &promptui.Select{
		Label:     "Which editor would you like to use?",
		IsVimMode: true,
		Items:     GetSupportedEditors(),
	}
	_, ed, err := editor_prompt.Run()
	return ed, err
}
