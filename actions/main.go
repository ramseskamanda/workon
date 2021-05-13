package actions

import (
	"os"
	"path/filepath"

	"github.com/ramseskamanda/workon/system"
	"github.com/urfave/cli/v2"
)

var (
	project          string
	editor           system.Editor
	editor_selection string
	latest           bool
	environment      system.Environment
)

func GetMainActionFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "project",
			Aliases:     []string{"proj", "p"},
			Value:       system.PROJECT_DEFAULT,
			Usage:       "the project to open",
			Destination: &project,
			Required:    false,
		},
		&cli.StringFlag{
			Name:        "editor",
			Aliases:     []string{"e"},
			Usage:       "editor to open the project into",
			Destination: &editor_selection,
			Required:    false,
		},
	}
}

func MainAction(c *cli.Context) error {
	environment = system.Environment(os.Getenv("TERM_PROGRAM"))
	var config system.Config
	var index system.InternalIndex
	err := system.LoadInternalIndex(&index)
	if err != nil {
		return err
	}
	if err := system.LoadConfig(&config); err != nil {
		err = system.AskInitializationQuestions(&config)
		if err != nil {
			return err
		}
		err = config.Save()
		if err != nil {
			return nil
		}
	}

	scanner := make(chan error)
	go system.Scan(config.Path, scanner, &index)

	if editor_selection != "" {
		editor = system.Editor(editor_selection)
	} else {
		editor = system.Editor(config.Editor)
	}

	if project == system.PROJECT_DEFAULT {
		project, err = system.AskForProject(&index)
		if err != nil {
			return err
		}

		if project == system.NEW_PROJECT_PROMPT {
			project, err = system.StartNewProject(&config, &index)
			if err != nil {
				return err
			}
		}

		if config.Editor == "" {
			editor_choice, err := system.AskForEditor()
			if err != nil {
				return err
			}
			editor = system.Editor(editor_choice)
			config.Editor = editor_choice
			config.Save()
		}
	}

	err = <-scanner
	if err != nil {
		return err
	}
	project, err = index.FindProjectByName(project)
	if err != nil {
		return err
	}

	index.MostRecent = append(index.MostRecent, filepath.Base(project))
	err = index.Overwrite()
	if err != nil {
		return err
	}

	return editor.Open(project)
}
