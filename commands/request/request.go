package request

import (
	"fmt"
	"github.com/adel-hadadi/tisan/constant"
	"github.com/iancoleman/strcase"
	"html/template"
	"log"
	"os"
	"strings"
)

type RequestCommand struct {
}

const (
	EnterRequestName = "please enter request name \n \t example: request:make [REQUEST_NAME]"
)

func (c *RequestCommand) Make(args []string) {
	if len(args) == 0 {
		fmt.Println(EnterRequestName)
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
		"Request": strcase.ToCamel(fName),
	}

	for i, a := range args[1:] {
		if a == "--model" {
			log.Println(a[i+1])
			data["IsResource"] = true
		}
	}

	tmpl, err := template.ParseFiles(constant.TemplateDirectory + constant.RequestTemplate)
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

func (c *RequestCommand) Help(args []string) {

}
