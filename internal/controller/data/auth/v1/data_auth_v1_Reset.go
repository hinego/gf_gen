package data

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	authV1 "github.com/sucold/starter/api/data/v1/auth"
)

func (r *controllerV1) Reset(ctx context.Context, req *authV1.ResetReq) (res *authV1.ResetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
