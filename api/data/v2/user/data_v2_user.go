package v2

import (
	"context"
)


type IUser interface {
	Info(ctx context.Context, req *InfoReq) (res *InfoRes, err error)
	Register(ctx context.Context, req *RegisterReq) (res *RegisterRes, err error)
	Reset(ctx context.Context, req *ResetReq) (res *ResetRes, err error)
	Logout(ctx context.Context, req *LogoutReq) (res *LogoutRes, err error)
	
}
