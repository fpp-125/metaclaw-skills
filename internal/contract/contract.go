package contract

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Contract struct {
	APIVersion    string        `yaml:"apiVersion"`
	Kind          string        `yaml:"kind"`
	Metadata      Metadata      `yaml:"metadata"`
	Permissions   Permissions   `yaml:"permissions"`
	Compatibility Compatibility `yaml:"compatibility"`
}

type Metadata struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Permissions struct {
	Network string   `yaml:"network"`
	Env     []string `yaml:"env"`
	Secrets []string `yaml:"secrets"`
}

type Compatibility struct {
	RuntimeTargets []string `yaml:"runtimeTargets"`
}

func Load(path string) (Contract, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Contract{}, err
	}
	var c Contract
	if err := yaml.Unmarshal(b, &c); err != nil {
		return Contract{}, err
	}
	if err := Validate(c); err != nil {
		return Contract{}, err
	}
	return c, nil
}

func Validate(c Contract) error {
	if strings.TrimSpace(c.APIVersion) != "metaclaw.capability/v1" {
		return fmt.Errorf("apiVersion must be metaclaw.capability/v1")
	}
	if strings.TrimSpace(c.Kind) != "CapabilityContract" {
		return fmt.Errorf("kind must be CapabilityContract")
	}
	if strings.TrimSpace(c.Metadata.Name) == "" {
		return fmt.Errorf("metadata.name is required")
	}
	if strings.TrimSpace(c.Metadata.Version) == "" {
		return fmt.Errorf("metadata.version is required")
	}
	switch strings.TrimSpace(c.Permissions.Network) {
	case "none", "outbound", "all":
	default:
		return fmt.Errorf("permissions.network must be one of none|outbound|all")
	}
	if len(c.Compatibility.RuntimeTargets) == 0 {
		return fmt.Errorf("compatibility.runtimeTargets must not be empty")
	}
	for _, rt := range c.Compatibility.RuntimeTargets {
		rt = strings.TrimSpace(rt)
		switch rt {
		case "podman", "docker", "apple_container":
		default:
			return fmt.Errorf("unsupported runtime target: %s", rt)
		}
	}
	return nil
}
