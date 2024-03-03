package config

import (
	"fmt"
	"slices"
)

type Config struct {
	TemplateDirectory string
	Framework         string
}

type Template struct {
	Directory string
	Handler   string
}

// ErrNotAllowedFramework error message that occurred when chosen framework not accepted
var ErrNotAllowedFramework = "%s is not allowed framework"

// Frameworks list of supported frameworks
var Frameworks = []string{
	"none",
	"echo",
	"gin",
	"fiber",
}

func NewConfig(framework string) (*Config, error) {
	if framework != "none" && !slices.Contains(Frameworks, framework) {
		return &Config{}, fmt.Errorf(ErrNotAllowedFramework, framework)
	}

	return &Config{
		TemplateDirectory: "./templates",
		Framework:         framework,
	}, nil
}
