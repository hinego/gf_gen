
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LogoutReq struct {
	g.Meta `path:"/v1/user/auth/logout" tags:"Auth" method:"post" summary:""`
}
type LogoutRes struct {
	g.Meta `mime:"application/json" example:""`
}