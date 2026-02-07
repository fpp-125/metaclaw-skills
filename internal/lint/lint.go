package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/metaclaw/metaclaw-skills/internal/contract"
)

type Result struct {
	SkillDir string
	Issues   []string
}

func Run(root string) ([]Result, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	results := make([]Result, 0)
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		skillDir := filepath.Join(root, e.Name())
		issues := checkSkill(skillDir)
		results = append(results, Result{SkillDir: skillDir, Issues: issues})
	}
	return results, nil
}

func checkSkill(skillDir string) []string {
	issues := make([]string, 0)
	skillPath := filepath.Join(skillDir, "SKILL.md")
	if _, err := os.Stat(skillPath); err != nil {
		issues = append(issues, "missing SKILL.md")
	}
	contractPath := filepath.Join(skillDir, "capability.contract.yaml")
	if _, err := os.Stat(contractPath); err != nil {
		issues = append(issues, "missing capability.contract.yaml")
	} else {
		if _, err := contract.Load(contractPath); err != nil {
			issues = append(issues, fmt.Sprintf("invalid capability contract: %v", err))
		}
	}
	if strings.Contains(filepath.Base(skillDir), " ") {
		issues = append(issues, "skill directory should not contain spaces")
	}
	return issues
}
