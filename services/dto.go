package services

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"html/template"
	"log"
	"os"
	"strings"
)

type DTOCommand struct {
	Destination string
	Options     HandlerOptions
}

type DTOOptions struct {
}

func (h *DTOCommand) Make() {
	nestedDir := strings.Split(h.Destination, "/")

	fName := nestedDir[len(nestedDir)-1]
	if len(nestedDir) > 1 {
		nestedDir = nestedDir[:len(nestedDir)-1]
		err := os.MkdirAll(strings.Join(nestedDir, "/"), os.ModePerm)
		if err != nil {
			fmt.Println("error according create folders")
			return
		}
	}

	data := map[string]interface{}{
		"Name": strcase.ToCamel(fName),
	}

	tmpl, err := template.ParseFiles(templatesBasePath + "dto.tmpl")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	file, err := os.Create(fmt.Sprintf("%s.go", h.Destination))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("DTO created successfully 👍")
}
