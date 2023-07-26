
package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LogoutReq struct {
	g.Meta `path:"/v2/data/user/logout" tags:"User" method:"post" summary:""`
}
type LogoutRes struct {
	g.Meta `mime:"application/json" example:""`
}
