package main

import (
	"log"
	"os"

	"github.com/ramseskamanda/workon/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags:  actions.GetMainActionFlags(),
		Action: actions.MainAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
