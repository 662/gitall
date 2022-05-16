package commands

import (
	"fmt"
	"gitall/pkg/services"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

func GitallAction(c *cli.Context) error {
	start := time.Now()

	var args = c.Args().Slice()
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := ExecGitCommandInSubdirectory(dir, args); err != nil {
		return err
	}
	elapsed := time.Since(start)

	log.Printf("completed in %s", elapsed)
	return nil
}

func ExecGitCommandInSubdirectory(basedir string, args []string) error {
	red := color.FgRed.Render
	files, err := ioutil.ReadDir(basedir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		var dir = path.Join(basedir, file.Name())

		log.Printf("run command \"git %s\" in %s", strings.Join(args[:], " "), dir)

		if msg, err := services.ExecGitCommand(dir, args...); err != nil {
			fmt.Printf("%s\n", red(msg))
		}
	}
	return nil
}
