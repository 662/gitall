package commands

import (
	"errors"
	"gitall/pkg/services"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var CloneCommand = cli.Command{
	Name:   "clone",
	Usage:  "Multiple clone",
	Flags:  CloneFlags,
	Action: CloneAction,
}

var CloneFlags = []cli.Flag{
	&cli.StringFlag{Name: "scm", Usage: "SCM type", Required: true},
	&cli.StringFlag{Name: "base-url", Usage: "SCM base url", Required: true},
	&cli.StringFlag{Name: "group", Usage: "SCM group", Required: true},
	&cli.StringFlag{Name: "token", Usage: "SCM token", Required: true},
}

func CloneAction(c *cli.Context) error {
	start := time.Now()

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	urls, err := getRepositoryUrls(c.String("scm"), c.String("group"), c.String("base-url"), c.String("token"))
	if err != nil {
		return err
	}
	if err := GitCloneAll(dir, urls); err != nil {
		return err
	}

	elapsed := time.Since(start)
	log.Printf("completed in %s", elapsed)
	return nil
}

func getRepositoryUrls(scm string, groupName string, baseUrl string, token string) (urls []string, err error) {
	switch scm {
	case "gitlab":
		return services.FetchGitlabRepositoriesByGroupName(baseUrl, groupName, token)
	default:
		return nil, errors.New("unsupported SCM type")
	}
}

func GitCloneAll(dir string, urls []string) error {
	for _, url := range urls {
		log.Printf("cloneing %s", url)
		if _, err := services.ExecGitCommand(dir, "clone", url); err != nil {
			return err
		}
	}
	return nil
}
