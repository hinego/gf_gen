package data

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	authV2 "github.com/sucold/starter/api/data/v2/auth"
)

func (r *controllerV2) Reset(ctx context.Context, req *authV2.ResetReq) (res *authV2.ResetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
