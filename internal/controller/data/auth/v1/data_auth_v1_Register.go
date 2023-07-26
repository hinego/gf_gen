package data

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	authV1 "github.com/sucold/starter/api/data/v1/auth"
)

func (r *controllerV1) Register(ctx context.Context, req *authV1.RegisterReq) (res *authV1.RegisterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
