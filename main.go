package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
)

type issueConfig struct {
	ownerName  string
	repoName   string
	issueLabel string
}

func parseArgs(args []string) *issueConfig {
	return &issueConfig{
		ownerName:  string(args[1]),
		repoName:   string(args[2]),
		issueLabel: string(args[3]),
	}
}

func main() {
	fmt.Println("Format: fetchin -[owner name] -[repo name] -[issue label]")

	info := os.Args

	issueInfo := parseArgs(info)

	fmt.Println(issueInfo.repoName)

	ctx := context.Background()
	client := github.NewClient(nil)

	issList, _, _ := client.Issues.ListByRepo(ctx, issueInfo.ownerName, issueInfo.repoName, nil)
	//val := issList[0]
	for _, issue := range issList {
		//fmt.Println(val, "\n\n=====================================\n\n")
		for _, label := range issue.Labels {
			if label.GetName() == issueInfo.issueLabel {
				fmt.Println(issue.GetNumber(), "\t\t", issue.GetURL())
			}
		}

	}
}
