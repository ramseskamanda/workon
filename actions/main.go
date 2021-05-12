package actions

import (
	"os"

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
	if err := system.LoadConfig(&config); err != nil {
		err = system.AskInitializationQuestions(&config)
		if err != nil {
			return err
		}
		err = config.Create()
		if err != nil {
			return nil
		}
	}

	channel := make(chan error)
	go system.Scan(config.Path, channel)

	if editor_selection != "" {
		editor = system.Editor(editor_selection)
	} else {
		editor = system.Editor(config.Editor)
	}

	if project == system.PROJECT_DEFAULT {
		//TODO: rewrite this part next
		project, editor = system.AskQuestions(latest)
	}

	err := <-channel
	index := system.InternalIndex{}
	project, err := index.FindProjectByName(project)
	if err != nil {
		return err
	}

	return editor.Open(config.Path + project)
}
