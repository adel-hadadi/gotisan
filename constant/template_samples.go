package constant

const (
	HandlerSample = `package handlers

import (
{{ if .IsRestful }}
"net/http"
{{ end }}
)

type {{ .Name }}Handler struct {

}

func (c *{{ .Name}}Handler) Init{{ .Name}}Handler() {

}

{{ if .IsRestful }}

func (c *{{ .Name}}Handler) Index(w http.ResponseWriter, r *http.Request) {

}


func (c *{{ .Name}}Handler) Store(w http.ResponseWriter, r *http.Request) {

}

func (c *{{ .Name}}Handler) Update(w http.ResponseWriter, r *http.Request) {

}


func (c *{{ .Name}}Handler) Delete(w http.ResponseWriter, r *http.Request) {

}

{{ end }}`

	ModelSample = `package models

import (
    "time"
)

type {{.Model}} struct {
    ID uint
    CreatedAt time.Duration
    UpdatedAt time.Duration
    DeletedAt time.Duration
}`
)
