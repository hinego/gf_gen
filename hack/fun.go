package main

import (
	"fmt"
)

const interfaceTemplate = `
// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package {{.ApiName}}

import (
{{range $version := .Versions}}
{{range .Data}}
	{{.Name}}{{$version.Name | title}} "github.com/sucold/starter/api/{{$.ApiName}}/{{$version.Name}}/{{.Name}}"
{{- end}}
{{- end}}
)

type Interface interface { {{range .Versions}}
	{{title .Name}}() {{title .Name}}Interface
{{- end}}
}

{{range $version := .Versions}}
type {{$version.Name | title}}Interface interface { {{range .Data}}
	{{.Name | title}}() {{.Name}}{{$version.Name | title}}.I{{.Name | title}}
    {{- end}}
}
{{- end}}
`

func genInterface(api API) (err error) {
	var input = &Input{
		Code: interfaceTemplate,
		Data: map[string]interface{}{
			"ApiName":  api.Name,
			"Versions": api.Data,
		},
		File: fmt.Sprintf("./api/%s/%s.go", api.Name, api.Name),
		Must: true,
	}
	return execute(input) // ./api/{{api.name}}/{{api.name}}.go
}
