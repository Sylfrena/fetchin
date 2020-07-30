package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

type Dependency interface {
	Get(issueInfo *issueConfig) []*github.Issue
}

type issueConfig struct {
	ownerName string
	repoName  string
}

type githubService struct {
	client *github.Client
	ctx    context.Context
}

func newGithubService() *githubService {
	return &githubService{
		client: github.NewClient(nil),
		ctx:    context.Background(),
	}
}

func (github *githubService) Get(issueInfo *issueConfig) []*github.Issue {

	issList, _, _ := github.client.Issues.ListByRepo(github.ctx, issueInfo.ownerName, issueInfo.repoName, nil)
	return issList
}

func parseArgs(args []string) []*issueConfig {

	issueSlice := make([]*issueConfig, len(args))

	for index, issueArgs := range args {
		argSplitty := strings.Split(issueArgs, ":")
		issueSlice[index] = &issueConfig{
			ownerName: string(strings.Trim(argSplitty[0], " ")),
			repoName:  string(strings.Trim(argSplitty[1], " ")),
		}
	}

	for _, j := range issueSlice {
		fmt.Println(*j)
	}

	return issueSlice
}

func getIssue(service Dependency, issueInfo *issueConfig, limit int, issueLabel string) {
	count := 1

	issList := service.Get(issueInfo)

	for _, issue := range issList {
		for _, label := range issue.Labels {
			if label.GetName() == issueLabel && count <= limit { //add clause if label != pr
				fmt.Println(issue.GetNumber(), "\t\t", issue.GetURL())
				count++
			}
			//			fmt.Println(label.GetName(), issue.GetURL())

		}
	}

}

func main() {

	//fmt.Println("Format: fetchin -[owner name] -[repo name] -[issue label]")
	gitService := newGithubService()

	limit := flag.Int("limit", 10, "maximum number of results")
	label := flag.String("label", "bug: major", "issue label")
	flag.Parse()
	//label := "bug: minor"
	info := flag.Args()

	//issueInfo := make([]*issueConfig, len(info))
	issueInfo := parseArgs(info)
	labelParse := strings.Split(*label, ",")

	//support for multiple labels
	for _, issueDetails := range issueInfo {
		fmt.Println("\n\n\n=============================================================================================\n\nFor ", issueDetails)
		for _, label := range labelParse {
			fmt.Println("\n----------------------------------------------------------------------------------------\n\nFor label = ", strings.Trim(label, " "))
			getIssue(gitService, issueDetails, *limit, strings.Trim(label, " "))
		}

	}
}
