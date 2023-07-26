package data

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	userV2 "github.com/sucold/starter/api/data/v2/user"
)

func (r *controllerV2) Register(ctx context.Context, req *userV2.RegisterReq) (res *userV2.RegisterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
