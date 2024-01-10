package services

import (
	"fmt"
	"github.com/adel-hadadi/gotisan/config"
	"github.com/adel-hadadi/gotisan/constant"
	"github.com/iancoleman/strcase"
	"html/template"
	"log"
	"strings"
)

type ModelCommand struct {
	Destination  string
	ModelOptions ModelOptions
}

type ModelOptions struct {
}

func (c *ModelCommand) Make(cfg *config.Config) {
	if checkFileExists(c.Destination) {
		fmt.Println(constant.ErrFileAlreadyExists)
		return
	}

	path := strings.Split(c.Destination, "/")

	if len(path) > 1 {
		err := createNestedFolder(path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	fName := path[len(path)-1]
	var data map[string]interface{} = map[string]interface{}{
		"Model": strcase.ToCamel(fName),
	}

	tmpl, err := template.ParseFiles(templatesBasePath + "model.tmpl")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	file, err := MakeFile(c.Destination)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("model create successfully :)")
}
