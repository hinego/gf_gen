package main

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/sucold/starter/internal/controller"
	_ "github.com/sucold/starter/internal/packed"
	"reflect"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		var ctx = context.Background()

		controller.User()
		controller.User().V1()
		controller.User().V1().Auth()
		controller.User().V1().Auth().Register(ctx, nil)
		/*
			gf gen api /api_name/v1/name/action
			还可以做一个这样的命令，这样就可以直接新增一个方法，在
			/api/api_name/v1/name/ 目录下自动生成 ActionReq的和ActionRes
			/controller/api_name/name/v1/ 目录下自动生成 func(*XX) Action(ctx .....)

		*/
		//SDK生成后也可以用类似上面的写法访问接口，例如
		//sdk..User().V1().Auth().Register(ctx, &InputRes{})

		//然后 *ghttp.RouterGroupd 的 group.Bind 可以支持一下
		//类似这样的写法
		//group.Bind(controller.Data())
		//group.Bind(controller.User().V1())
		//group.Bind(controller.User().V1().Auth())
		group.Bind(Convent(controller.Data())...)
	})
	s.SetPort(658)
	s.Run()
}
func Convent(data any) []any {
	h := &handler{}
	h.getAllMethods(reflect.ValueOf(data))
	return h.data
}

type handler struct {
	data []any
}

func (r *handler) getAllMethods(dataValue reflect.Value) {
	if r.checkAll(dataValue) == nil {
		r.data = append(r.data, dataValue.Interface())
		return
	}
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumMethod(); i++ {
		if dataValue.Method(i).Type().NumIn() != 0 || dataValue.Method(i).Type().NumOut() != 1 {
			continue
		}
		results := dataValue.Method(i).Call(nil)
		for _, result := range results {
			r.getAllMethods(result)
		}
	}
	return
}
func (r *handler) checkAll(dataValue reflect.Value) (err error) {
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumMethod(); i++ {
		method := dataType.Method(i)
		if err = r.check(method.Type); err != nil {
			return
		}
	}
	return
}
func (r *handler) check(reflectType reflect.Type) (err error) {
	if reflectType.NumIn() != 2 || reflectType.NumOut() != 2 {
		return gerror.New("invalid handler")
	}

	if !reflectType.In(0).Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		return gerror.New("invalid handler")
	}

	if !reflectType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return gerror.New("invalid handler")
	}

	// The request struct should be named as `xxxReq`.
	if !gstr.HasSuffix(reflectType.In(1).String(), `Req`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for request: defined as "%s", but it should be named with "Req" suffix like "XxxReq"`,
			reflectType.In(1).String(),
		)
		return
	}

	// The response struct should be named as `xxxRes`.
	if !gstr.HasSuffix(reflectType.Out(0).String(), `Res`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for response: defined as "%s", but it should be named with "Res" suffix like "XxxRes"`,
			reflectType.Out(0).String(),
		)
		return
	}
	return
}
