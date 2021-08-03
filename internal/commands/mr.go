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

	"github.com/urfave/cli/v2"
)

var MRCommand = cli.Command{
	Name:   "mr",
	Usage:  "Multiple create gitlab merge request",
	Flags:  MRFlags,
	Action: MRAction,
}

var MRFlags = []cli.Flag{
	&cli.StringFlag{Name: "base-url", Usage: "SCM base url", Required: true},
	&cli.StringFlag{Name: "token", Usage: "SCM token", Required: true},
}

func MRAction(c *cli.Context) error {
	start := time.Now()

	args := c.Args().Slice()
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	baseUrl := c.String("base-url")
	token := c.String("token")
	if len(args) < 2 {
		log.Fatalf("[error]\tMust has source_branch and trarget_branh.")
	}
	sourceBranch := args[0]
	trargetBranh := args[1]

	names, err := getProjectNames(dir)
	if err != nil {
		log.Fatalf("[error]\tgetProjectNames error.")
		return err
	}
	log.Printf("Projecct names %s", names)
	projects, err := getProjectIds(names, baseUrl, token)
	if err != nil {
		log.Fatalf("[error]\tgetProjectIds error.")
		return err
	}
	createMergeRequest(projects, baseUrl, token, sourceBranch, trargetBranh)

	elapsed := time.Since(start)
	log.Printf("[info]\tcompleted in %s", elapsed)
	return nil
}

func getProjectNames(basedir string) ([]string, error) {
	files, err := ioutil.ReadDir(basedir)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(files))
	for i, file := range files {
		if !file.IsDir() {
			continue
		}
		var dir = path.Join(basedir, file.Name())
		msg, err := services.ExecGitCommand(dir, "ls-remote", "--get-url")
		if err != nil {
			log.Printf("[info]\tExecGitCommand 'git ls-remote --get-url' at %s error.", dir)
			log.Fatalln(msg)
			return nil, err
		}
		msg = strings.TrimRight(msg, "\000")
		msg = strings.Replace(msg, "\n", "", -1)
		name := strings.Replace(strings.Split(msg, ":")[1], ".git", "", -1)
		names[i] = name
	}
	return names, nil
}

func getProjectIds(names []string, baseUrl string, token string) (map[string]string, error) {
	projects := make(map[string]string)
	for _, name := range names {
		log.Printf("[info]\tGet project %s id", name)
		project, err := services.FindGitlabProject(baseUrl, name, token)
		if err != nil {
			return nil, err
		}
		projects[name] = fmt.Sprintf("%d", int(project["id"].(float64)))
	}
	return projects, nil
}

func createMergeRequest(projects map[string]string, baseUrl string, token string, sourceBranch string, targetBranch string) {
	for key := range projects {
		log.Printf("[info]\tCreate MR on %s..", key)
		if _, err := services.CreateGitlabMergeRequest(baseUrl, projects[key], token, sourceBranch, targetBranch); err != nil {
			log.Printf("%s", err)
		}
		log.Println("[info]\tok!")
	}
}
