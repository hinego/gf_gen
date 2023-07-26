
package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ResetReq struct {
	g.Meta `path:"/v2/data/user/reset" tags:"User" method:"post" summary:""`
}
type ResetRes struct {
	g.Meta `mime:"application/json" example:""`
}
