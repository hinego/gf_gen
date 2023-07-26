package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"strings"
)

func generateStructFile(api API, version Version, file File) (err error) {
	for _, fun := range file.Data {
		fun.Name = cases.Title(language.English).String(fun.Name)
		var input = &Input{
			Code: structTemplate,
			Data: &StructInput{
				Function:    fun,
				VersionName: version.Name,
			},
			File: fmt.Sprintf("./api/%s/%s/%s/%s_%s_%s_%s.go", api.Name, version.Name, file.Name, api.Name, version.Name, file.Name, fun.Name),
		}
		if err = execute(input); err != nil { // ./api/{{api.name}}/{{version.name}}/{{file.name}}/{{api.name}}_{{version.name}}_{{file.name}}_{{fun.name}}.go
			return
		}
	}
	var input = &Input{
		Code: structInterface,
		Data: &StructInterfaceInput{
			Functions:   file.Data,
			VersionName: version.Name,
			FileName:    strings.Title(file.Name),
		},
		File: fmt.Sprintf("./api/%s/%s/%s/%s_%s_%s.go", api.Name, version.Name, file.Name, api.Name, version.Name, file.Name),
		Must: true,
	}
	return execute(input) // ./api/{{api.name}}/{{version.name}}/{{file.name}}/{{api.name}}_{{version.name}}_{{file.name}}.go
}
func generateInitFile(api API, version Version, file File) error {
	var input = &Input{
		Code: initTemplate,
		Data: map[string]interface{}{
			"ApiName":     api.Name,
			"FileName":    file.Name,
			"VersionName": version.Name,
		},
		File: fmt.Sprintf("./internal/controller/%s/%s/%s/%s_%s_%s.go", api.Name, file.Name, version.Name, api.Name, file.Name, version.Name),
	}
	return execute(input) // ./internal/controller/{{api.name}}/{{file.name}}/{{version.name}}/{{file.name}}_{{version.name}}_init.go
}
func generateControllerFile(api API, version Version, file File, function Function) error {
	var input = &Input{
		Code: controllerTemplate,
		Data: map[string]any{
			"ApiName":      api.Name,
			"VersionName":  version.Name,
			"FileName":     file.Name,
			"FunctionName": function.Name,
		},
		File: fmt.Sprintf("./internal/controller/%s/%s/%s/%s_%s_%s_%s.go", api.Name, file.Name, version.Name, api.Name, file.Name, version.Name, function.Name),
	}
	return execute(input)
}
func main() {
	var apis = append(data, admin)
	apis = append(apis, admin2)
	err := generateControllerPackage(apis)
	if err != nil {
		log.Printf("error generating controller package: %v", err)
		return
	}
	for _, api := range apis {
		if err = genInterface(api); err != nil {
			return
		}
		for _, version := range api.Data {
			for _, file := range version.Data {
				for _, fun := range file.Data {
					if fun.Path == "" {
						fun.Path = fmt.Sprintf("/%s/%s/%s/%s", version.Name, api.Name, file.Name, fun.Name)
					}
					if fun.Tags == "" {
						fun.Tags = strings.Title(file.Name)
					}
					if fun.Method == "" {
						fun.Method = "post"
					}
					if fun.Mime == "" {
						fun.Mime = "application/json"
					}
				}
				err := generateStructFile(api, version, file)
				if err != nil {
					fmt.Println(err)
					return
				}
				err = generateInitFile(api, version, file)
				if err != nil {
					fmt.Println(err)
					return
				}
				for _, function := range file.Data {
					err := generateControllerFile(api, version, file, *function)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
}
