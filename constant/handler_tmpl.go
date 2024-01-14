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

func New{{ .Name}}Handler() *{{ .Name }}Handler {
	return &{{ .Name }}Handler{}
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

	GinHandlerSample = `package handlers

import (
{{ if .IsRestful }}
"github.com/gin-gonic/gin"
"net/http"
{{ end }}
)

type {{ .Name }}Handler struct {

}

func New{{ .Name}}Handler() *{{ .Name }}Handler {
	return &{{ .Name }}Handler{}
}

{{ if .IsRestful }}

func (c *{{ .Name}}Handler) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOk, gin.H{})
}


func (c *{{ .Name}}Handler) Store(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{})
}

func (c *{{ .Name}}Handler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOk, gin.H{})
}


func (c *{{ .Name}}Handler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOk, gin.H{})
}

{{ end }}`

	EchoHandlerSample = `package handlers

import (
{{ if .IsRestful }}
"github.com/labstack/echo/v4"
"net/http"
{{ end }}
)

type {{ .Name }}Handler struct {

}

func New{{ .Name}}Handler() *{{ .Name }}Handler {
	return &{{ .Name }}Handler{}
}

{{ if .IsRestful }}

func (c *{{ .Name}}Handler) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOk, interface{})
}


func (c *{{ .Name}}Handler) Store(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, interface{})
}

func (c *{{ .Name}}Handler) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOk, interface{})
}


func (c *{{ .Name}}Handler) Delete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOk, interface{})
}

{{ end }}`

	FiberHandlerSample = `package handlers

import (
{{ if .IsRestful }}
"github.com/gofiber/fiber/v2"
"net/http"
{{ end }}
)

type {{ .Name }}Handler struct {

}

func New{{ .Name}}Handler() *{{ .Name }}Handler {
	return &{{ .Name }}Handler{}
}

{{ if .IsRestful }}

func (c *{{ .Name}}Handler) Index(ctx *fiber.Ctx) error {
	 return ctx.JSON(fiber.Map{})
}


func (c *{{ .Name}}Handler) Store(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{})
}

func (c *{{ .Name}}Handler) Update(ctx *fiber.Ctx) error {
	 return ctx.JSON(fiber.Map{})
}


func (c *{{ .Name}}Handler) Delete(ctx *fiber.Ctx) error {
	 return ctx.JSON(fiber.Map{})
}

{{ end }}`
)
