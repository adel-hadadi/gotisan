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

var AllowedFrameworks = []string{
	"echo",
	"gin",
	"fiber",
}

func NewDefaultConfig(framework string) (*Config, error) {
	if framework != "" && !slices.Contains(AllowedFrameworks, framework) {
		return &Config{}, fmt.Errorf(ErrNotAllowedFramework, framework)
	}

	return &Config{
		TemplateDirectory: "./templates",
		Framework:         framework,
	}, nil
}
