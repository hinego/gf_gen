package main

const controllerPackageTemplate = `package controller

import (
{{- range $api := .APIs}}
	{{$api.Name}}Face "github.com/sucold/starter/api/{{$api.Name}}"
{{- range $version := .Data}}{{- range $file := .Data}}
	{{$api.Name}}{{title $file.Name}}{{title $version.Name}}Api "github.com/sucold/starter/api/{{$api.Name}}/{{$version.Name}}/{{$file.Name}}"
{{- end}}{{- end}}{{- end}}
)

var (
{{- range $api := .APIs}}{{- range $version := .Data}}{{- range $file := .Data}}
	_{{$api.Name}}{{title $file.Name}}{{title $version.Name}} {{$api.Name}}{{title $file.Name}}{{title $version.Name}}Api.I{{title $file.Name}}
{{- end}}{{- end}}{{- end}}
)

func Register(name string, data any) {
	var ok bool
	switch name {
{{- range $api := .APIs}}{{- range $version := .Data}}{{- range $file := .Data}}
	case "{{$api.Name}}{{title $file.Name}}{{title $version.Name}}":
		if _{{$api.Name}}{{title $file.Name}}{{title $version.Name}},ok = data.({{$api.Name}}{{title $file.Name}}{{title $version.Name}}Api.I{{title $file.Name}});!ok {
			panic("{{$api.Name}}{{title $file.Name}}{{title $version.Name}} register error")
		}
		break
{{- end}}{{- end}}{{- end}}
	}
}

{{- range $api := .APIs}}
type {{$api.Name}} struct{}
func {{$api.Name | title}}() {{$api.Name}}Face.Interface {
	return &{{$api.Name}}{}
}
{{- end}}
{{- range $api := .APIs}}{{- range $version := .Data}}
func (r *{{$api.Name}}) {{title $version.Name}}() {{$api.Name}}Face.{{title $version.Name}}Interface {
	return &{{$api.Name}}{{title $version.Name}}{}
}
{{- end}}{{- end}}
{{- range $api := .APIs}}{{- range $version := .Data}}
type {{$api.Name}}{{title $version.Name}} struct{}
{{- range $file := .Data}}
func (d *{{$api.Name}}{{title $version.Name}}) {{title $file.Name}}() {{$api.Name}}{{title $file.Name}}{{title $version.Name}}Api.I{{title $file.Name}} {
	return _{{$api.Name}}{{title $file.Name}}{{title $version.Name}}
}
{{- end}}{{- end}}{{- end}}
`
const importControllerTemplate = `package packed

import (
{{- range $api := .APIs}}{{- range $version := .Data}}{{- range $file := .Data}}
	_ "github.com/sucold/starter/internal/controller/{{$api.Name}}/{{$file.Name}}/{{$version.Name}}"
{{- end}}{{- end}}{{- end}}
)
`

func generateControllerPackage(apis []API) (err error) {
	if err = generateImportControllerPackage(apis); err != nil {
		return err
	}
	var input = &Input{
		Code: controllerPackageTemplate,
		Data: map[string]any{
			"APIs": apis,
		},
		File: "./internal/controller/controller.go",
		Must: true,
	}
	return execute(input) // ./internal/controller/controller.go
}
func generateImportControllerPackage(apis []API) error {
	var input = &Input{
		Code: importControllerTemplate,
		Data: map[string]any{
			"APIs": apis,
		},
		File: "./internal/packed/init_controller.go",
		Must: true,
	}
	return execute(input) // ./internal/controller/controller.go
}
