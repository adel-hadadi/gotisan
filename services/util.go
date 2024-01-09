package services

import (
	"errors"
	"fmt"
	"github.com/adel-hadadi/gotisan/constant"
	"os"
	"strings"
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

func checkFileExists(path string) bool {
	_, err := os.Stat(path + ".go")
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func createNestedFolder(path []string) error {
	nestedDir := path[:len(path)-1]
	err := os.MkdirAll(strings.Join(nestedDir, "/"), os.ModePerm)
	if err != nil {
		return errors.New(constant.ErrAccordingCreateFolder)
	}

	return nil
}
