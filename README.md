![gotisan](docs/gotisan-demo.png)

# Gotisan

gotisan is the command line interface inspired by laravel artisan for your golang projects that give you flexibility to have your custom file structures.


## Why should use this package?

- doing comfortable tasks such as creating new handler, model or DTO
- flexibility to have your own file structure with using template files

## Installation

you can easily install this package globally in your local machine and use it in your golang projects. for installing this package just run below command in your terminal.

```bash
  go install github.com/adel-hadadi/gotisan
```

## Usage
to start using this package in our project just should go in your root directory of project and first of all run `init` command.

``` bash
gotisan init 
```
this command will create a `.gotisan` folder in your root directory that contain template files. if you're using git in your project, gotisan will add .gotisan folder in .gitignore file.


#### Make Handler
``` bash
gotisan make:handler [name] [options]
```

by default above command will create a handler with this structure:
``` go
package handlers

type UserHandler struct {

}

func InitUserHandler() *UserHandler {

}
``` 
but if you want a custom handler structure just go in `.gotisan/templates/handler.tmp` and customize to what you want.

by passing `-r` or `--restful` to `make:handler` command, package will create a restful handler with CRUD methods.

``` go
package handlers

import (

"net/http"

)

type UserHandler struct {

}

func InitUserHandler() *UserHandler {

}



func (c *UserHandler) Index(w http.ResponseWriter, r *http.Request) {

}


func (c *UserHandler) Store(w http.ResponseWriter, r *http.Request) {

}

func (c *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

}


func (c *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

```

## Make Model
``` bash
gotisan make:model [name]
```
the above command just create a simple model with default structure that contain common use fields:

``` go
package models

import (
    "time"
)

type Car struct {
    ID uint
    CreatedAt time.Duration
    UpdatedAt time.Duration
    DeletedAt time.Duration
}
```

## Make DTO

``` bash 
gotisan make:dto [name]
```
make dto command will create request and response objects. below example is for profile dto:
```go
package dtos

type ProfileReq struct {
}

type ProfileRes struct {
}

func NewProfileReq(data interface{}) *ProfileReq {
	return &ProfileReq{}
}

func NewProfileRes(data interface{}) *ProfileRes {
	return &ProfileRes{}
}
```