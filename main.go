package main

import (
	"fmt"
	commands2 "github.com/adel-hadadi/tisan/commands"
	"github.com/adel-hadadi/tisan/commands/handler"
	"github.com/adel-hadadi/tisan/commands/model"
	"github.com/adel-hadadi/tisan/commands/request"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"reflect"
	"strings"
)

var commands map[string]commands2.Command = map[string]commands2.Command{
	"model":   &model.ModelCommand{},
	"handler": &handler.HandlerCommand{},
	"request": &request.RequestCommand{},
}

func main() {
	r := os.Args[1]
	structure := strings.Split(r, ":")

	object := reflect.ValueOf(commands[structure[0]])
	if !object.IsValid() {
		log.Fatalf("object %s is not valid", commands[structure[0]])
		return
	}
	fmt.Println(object)
	mName := strcase.ToCamel(strings.ToLower(structure[1]))
	method := object.MethodByName(mName)

	if !method.IsValid() {
		log.Fatalf("method %s not found", mName)
		return
	}

	fn, ok := method.Interface().(func([]string))
	if !ok {
		log.Fatal("something went wrong")
		return
	}

	fn(os.Args[2:])
}
