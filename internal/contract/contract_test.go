package contract

import "testing"

func TestValidateContract(t *testing.T) {
	valid := Contract{
		APIVersion: "metaclaw.capability/v1",
		Kind:       "CapabilityContract",
		Metadata: Metadata{
			Name:    "obsidian.search",
			Version: "v1.0.0",
		},
		Permissions:   Permissions{Network: "none"},
		Compatibility: Compatibility{RuntimeTargets: []string{"docker", "podman"}},
	}
	if err := Validate(valid); err != nil {
		t.Fatalf("validate valid: %v", err)
	}
	invalid := valid
	invalid.Permissions.Network = "internet"
	if err := Validate(invalid); err == nil {
		t.Fatal("expected invalid network error")
	}
}
