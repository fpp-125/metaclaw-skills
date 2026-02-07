package lint

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunLintsSkillDirectories(t *testing.T) {
	root := t.TempDir()
	valid := filepath.Join(root, "obsidian.search")
	if err := os.MkdirAll(valid, 0o755); err != nil {
		t.Fatalf("mkdir valid: %v", err)
	}
	if err := os.WriteFile(filepath.Join(valid, "SKILL.md"), []byte("# x\n"), 0o644); err != nil {
		t.Fatalf("write skill: %v", err)
	}
	contract := `apiVersion: metaclaw.capability/v1
kind: CapabilityContract
metadata:
  name: obsidian.search
  version: v1.0.0
permissions:
  network: none
compatibility:
  runtimeTargets: [docker]
`
	if err := os.WriteFile(filepath.Join(valid, "capability.contract.yaml"), []byte(contract), 0o644); err != nil {
		t.Fatalf("write contract: %v", err)
	}
	invalid := filepath.Join(root, "bad skill")
	if err := os.MkdirAll(invalid, 0o755); err != nil {
		t.Fatalf("mkdir invalid: %v", err)
	}
	if err := os.WriteFile(filepath.Join(invalid, "SKILL.md"), []byte("# y\n"), 0o644); err != nil {
		t.Fatalf("write invalid skill: %v", err)
	}

	results, err := Run(root)
	if err != nil {
		t.Fatalf("run lint: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	foundInvalid := false
	for _, r := range results {
		if len(r.Issues) > 0 {
			foundInvalid = true
		}
	}
	if !foundInvalid {
		t.Fatal("expected at least one invalid skill result")
	}
}
