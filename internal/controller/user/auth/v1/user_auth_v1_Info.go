package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	authV1 "github.com/sucold/starter/api/user/v1/auth"
)

func (r *controllerV1) Info(ctx context.Context, req *authV1.InfoReq) (res *authV1.InfoRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
