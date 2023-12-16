package services

import (
	"fmt"
	"os"
)

const templatesBasePath = ".gotisan/templates/"

func MakeFile(destination string) (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Create(fmt.Sprintf("%s.go", wd+"/"+destination))
	if err != nil {
		return nil, err
	}

	return file, nil
}
