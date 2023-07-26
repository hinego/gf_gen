package main

import (
	"html/template"
	"strings"
)

const structInterface = `package {{.VersionName}}

import (
	"context"
)


type I{{.FileName}} interface {
	{{range .Functions}}{{.Name}}(ctx context.Context, req *{{.Name}}Req) (res *{{.Name}}Res, err error)
	{{end}}
}
`

const structTemplate = `
package {{.VersionName}}

import (
	"github.com/gogf/gf/v2/frame/g"
)

type {{.Name}}Req struct {
	g.Meta ` + "`path:\"{{.Path}}\" tags:\"{{.Tags}}\" method:\"{{.Method}}\" summary:\"{{.Summary}}\"`" + `
}
type {{.Name}}Res struct {
	g.Meta ` + "`mime:\"{{.Mime}}\" example:\"{{.Default}}\"`" + `
}
`

const controllerTemplate = `package {{.ApiName}}

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	{{.FileName}}{{.VersionName | title}} "github.com/sucold/starter/api/{{.ApiName}}/{{.VersionName}}/{{.FileName}}"
)

func (r *controller{{title .VersionName}}) {{title .FunctionName}}(ctx context.Context, req *{{.FileName}}{{.VersionName | title}}.{{title .FunctionName}}Req) (res *{{.FileName}}{{.VersionName | title}}.{{title .FunctionName}}Res, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
`
const initTemplate = `package {{.ApiName}}

import (
	"github.com/sucold/starter/internal/controller"
)

func init() {
	controller.Register("{{.ApiName}}{{title .FileName}}{{title .VersionName}}", &controller{{title .VersionName}}{})
}

type controller{{title .VersionName}} struct{}
`

var (
	funcMap = template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
	}
)
