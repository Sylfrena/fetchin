package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

type issueConfig struct {
	ownerName string
	repoName  string
}

func parseArgs(args []string) *issueConfig {
	argSplit := strings.Split(args[0], ":")
	

	
	return &issueConfig{
		ownerName: string(argSplit[0]),
		repoName:  string(argSplit[1]),
	}
}

func getIssue(issueInfo *issueConfig, limit int, issueLabel string) {

	ctx := context.Background()
	client := github.NewClient(nil)
	count := 1

	issList, _, _ := client.Issues.ListByRepo(ctx, issueInfo.ownerName, issueInfo.repoName, nil)
	//val := issList[0]
	for _, issue := range issList {
		for _, label := range issue.Labels {
			if label.GetName() == issueLabel && count <= limit {
				fmt.Println(issue.GetNumber(), "\t\t", issue.GetURL())
				count++
			}
		}
	}

}

func main() {

	//fmt.Println("Format: fetchin -[owner name] -[repo name] -[issue label]")

	limit := flag.Int("limit", 10, "maximum number of results")
	label := flag.String("label", "bug: major", "issue label")
	flag.Parse()
	//label := "bug: minor"
	info := flag.Args()

	//issueInfo := make([]*issueConfig, len(info))
	issueInfo := parseArgs(info)

	//support for multiple labels
	labelParse := strings.Split(*label, ",")
	for _, label := range labelParse {
		fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------\nFor label = ", strings.Trim(label, " "))
		getIssue(issueInfo, *limit, strings.Trim(label, " "))

	}

	//	fmt.Println("limit:", *limit, label_parse[1])
}
