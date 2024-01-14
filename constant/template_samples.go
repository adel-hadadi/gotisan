package constant

const (
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
	DTOSample = `package dtos

type {{ .Name }}Req struct {
}

type {{ .Name }}Res struct {
}

func New{{ .Name }}Req(data interface{}) *{{ .Name }}Req {
	return &{{ .Name }}Req{}
}

func New{{ .Name }}Res(data interface{}) *{{ .Name }}Res {
	return &{{ .Name }}Res{}
}`
)
