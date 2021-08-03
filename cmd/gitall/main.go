package main

import (
	"gitall/internal/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "GitAll",
		Usage:   "Git command for multiple repositories",
		Version: "0.0.1",
		Action:  commands.GitallAction,
		Commands: []*cli.Command{
			&commands.MRCommand,
			&commands.CloneCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
