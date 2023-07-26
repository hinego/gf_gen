
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta `path:"/v1/data/auth/register" tags:"Auth" method:"post" summary:""`
}
type RegisterRes struct {
	g.Meta `mime:"application/json" example:""`
}
