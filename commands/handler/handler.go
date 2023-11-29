package handler

import (
	"fmt"
	"github.com/adel-hadadi/tisan/constant"
	"github.com/iancoleman/strcase"
	"html/template"
	"log"
	"os"
	"strings"
)

type HandlerCommand struct {
}

const (
	EnterHandlerName = "please enter handler name \n \t example: handler:make [HANDLER_NAME]"
)

func (h *HandlerCommand) Make(args []string) {
	if len(args) == 0 {
		fmt.Println(EnterHandlerName)
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
	data := map[string]interface{}{
		"Name":       strcase.ToCamel(fName),
		"IsResource": false,
	}

	for _, a := range args[1:] {
		if a == "-r" || a == "--resource" {
			data["IsResource"] = true
		}
	}

	tmpl, err := template.ParseFiles(constant.TemplateDirectory + constant.HandlerTemplate)
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

func (h *HandlerCommand) Help(args []string) {
	file, err := os.ReadFile("templates/handler-help.tmpl")
	if err != nil {
		return
	}
	fmt.Println(string(file))
}
