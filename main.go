package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/go-github/github"
)

type issueConfig struct {
	ownerName  string
	repoName   string
	issueLabel string
}

func parseArgs(args []string) *issueConfig {
	return &issueConfig{
		ownerName:  string(args[0]),
		repoName:   string(args[1]),
		issueLabel: string(args[2]),
	}
}

func getIssue(issueInfo *issueConfig, limit int) {

	ctx := context.Background()
	client := github.NewClient(nil)
	count := 1

	issList, _, _ := client.Issues.ListByRepo(ctx, issueInfo.ownerName, issueInfo.repoName, nil)
    	//val := issList[0]
	for _, issue := range issList {
		for _, label := range issue.Labels {
			if label.GetName() == issueInfo.issueLabel && count <= limit {
				fmt.Println(issue.GetNumber(), "\t\t", issue.GetURL())
				count++
			}
		}
	}

}

func main() {

	//fmt.Println("Format: fetchin -[owner name] -[repo name] -[issue label]")

	limit := flag.Int("limit", 10, "maximum number of results")
	flag.Parse()

	info := flag.Args()

	issueInfo := parseArgs(info)

	getIssue(issueInfo, *limit)

	fmt.Println("limit:", *limit)
}
