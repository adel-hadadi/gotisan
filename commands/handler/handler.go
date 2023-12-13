package handler

import (
	"fmt"
	"github.com/adel-hadadi/tisan/config"
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

func (h *HandlerCommand) Make(path string, cfg *config.Config) {
	nestedDir := strings.Split(path, "/")

	if len(path) > 1 {
		nestedDir = nestedDir[:len(nestedDir)-1]
		err := os.MkdirAll(strings.Join(nestedDir, "/"), os.ModePerm)
		if err != nil {
			fmt.Println("error according create folders")
			return
		}
	}

	fName := nestedDir[len(nestedDir)-1]
	data := map[string]interface{}{
		"Name":       strcase.ToCamel(fName),
		"IsResource": false,
	}

	// todo: check args
	//for _, a := range args[1:] {
	//	if a == "-r" || a == "--resource" {
	//		data["IsResource"] = true
	//	}
	//}

	tmpl, err := template.ParseFiles(cfg.Template.Directory + cfg.Template.Handler)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	file, _ := os.Create(fmt.Sprintf("%s.go", fName))
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
