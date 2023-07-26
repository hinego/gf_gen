
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ResetReq struct {
	g.Meta `path:"/v1/data/auth/reset" tags:"Auth" method:"post" summary:""`
}
type ResetRes struct {
	g.Meta `mime:"application/json" example:""`
}
