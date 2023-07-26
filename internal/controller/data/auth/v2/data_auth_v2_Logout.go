package data

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	authV2 "github.com/sucold/starter/api/data/v2/auth"
)

func (r *controllerV2) Logout(ctx context.Context, req *authV2.LogoutReq) (res *authV2.LogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
