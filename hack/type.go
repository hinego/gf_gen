package main

import (
	"bytes"
	"github.com/gogf/gf/v2/os/gfile"
	"html/template"
	"log"
)

var data = make([]API, 0)

type API struct {
	Name string    //对应Hello目录
	Data []Version //对应Hello目录下的版本
}
type Version struct {
	Name string //对应v1目录
	Data []File //对应v1目录下的API文件，例如:user.go hello.go
}

type File struct {
	Name string      //对应user.go hello.go
	Data []*Function //对应user.go hello.go下的API
}

type Function struct {
	Name        string //对应 方法的名称 例如：HelloReq中的Hello
	Path        string //对应 方法的路径 例如：HelloReq中的/hello
	Tags        string //对应 方法的标签 例如：HelloReq中的hello
	Summary     string //对应 方法的简介 例如：HelloReq中的hello
	Method      string //对应 方法的请求方式 例如：HelloReq中的GET
	Deprecated  string //对应 方法的废弃 例如：HelloReq中的hello
	Description string //对应 方法的详细描述 例如：HelloReq中的hello
	Mime        string //对应 方法的MIME类型 例如：HelloReq中的text/html
	Type        string //对应 方法的类型 例如：HelloReq中的HelloRes
	In          string //对应 方法的提交方式 例如：HelloReq中的header/path/query/cookie
	Default     string //对应 方法的默认值 例如：HelloReq中的string
}
type StructInput struct {
	*Function
	VersionName string
}
type StructInterfaceInput struct {
	Functions   []*Function
	VersionName string
	FileName    string
}
type Input struct {
	Code string
	Data any
	File string
	Must bool
}

func execute(in *Input) (err error) {
	if gfile.Exists(in.File) && !in.Must {
		log.Println("skipfile", in.File)
		return nil
	}
	var code *template.Template
	if code, err = template.New("code").Funcs(funcMap).Parse(in.Code); err != nil {
		return
	}
	var buffer bytes.Buffer
	if err = code.Execute(&buffer, in.Data); err != nil {
		return
	}
	log.Println("generate", in.File)
	return gfile.PutContents(in.File, buffer.String())
}
