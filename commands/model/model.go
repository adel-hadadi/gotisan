package model

import (
	"fmt"
	"github.com/adel-hadadi/tisan/constant"
	"github.com/iancoleman/strcase"
	"html/template"
	"log"
	"os"
	"strings"
)

type ModelCommand struct {
}

const (
	EnterModelName = "please enter model name \n \t example: model:make [MODEL_NAME]"
)

func (c *ModelCommand) Make(args []string) {
	if len(args) == 0 {
		fmt.Println(EnterModelName)
		return
	}

	path := strings.Split(args[0], "/")
	if len(path) > 1 {
		nestedDir := path[:len(path)-1]
		err := os.MkdirAll(strings.Join(nestedDir, "/"), os.ModePerm)
		if err != nil {
			fmt.Println("error according create folders")
			return
		}
	}

	fName := path[len(path)-1]
	var data map[string]interface{} = map[string]interface{}{
		"Model": strcase.ToCamel(fName),
	}

	tmpl, err := template.ParseFiles(constant.TemplateDirectory + constant.ModelTemplate)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	file, _ := os.Create(fmt.Sprintf("%s.go", args[0]))
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func (c *ModelCommand) Help(args []string) {

}
