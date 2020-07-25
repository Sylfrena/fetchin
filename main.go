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

func getIssue(issueInfo *issueConfig, limit int, issueLabel string) {

	ctx := context.Background()
	client := github.NewClient(nil)
	count := 1

	issList, _, _ := client.Issues.ListByRepo(ctx, issueInfo.ownerName, issueInfo.repoName, nil)

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
	labelParse := strings.Split(*label, ",")

	//support for multiple labels
	for _, issueDetails := range issueInfo {
		fmt.Println("\n\n\n=============================================================================================\n\nFor ", issueDetails)
		for _, label := range labelParse {
			fmt.Println("\n----------------------------------------------------------------------------------------\n\nFor label = ", strings.Trim(label, " "))
			getIssue(issueInfo[0], *limit, strings.Trim(label, " "))
		}

	}
}
