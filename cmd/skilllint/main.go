package main

import (
	"fmt"
	"os"

	"github.com/fpp-125/metaclaw-skills/internal/lint"
)

func main() {
	root := "./skills"
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	results, err := lint.Run(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "skilllint failed: %v\n", err)
		os.Exit(1)
	}
	hasIssues := false
	for _, res := range results {
		if len(res.Issues) == 0 {
			fmt.Printf("OK  %s\n", res.SkillDir)
			continue
		}
		hasIssues = true
		fmt.Printf("ERR %s\n", res.SkillDir)
		for _, issue := range res.Issues {
			fmt.Printf("  - %s\n", issue)
		}
	}
	if hasIssues {
		os.Exit(1)
	}
}
