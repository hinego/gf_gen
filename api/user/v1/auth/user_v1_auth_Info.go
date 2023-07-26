
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type InfoReq struct {
	g.Meta `path:"/v1/user/auth/info" tags:"Auth" method:"post" summary:""`
}
type InfoRes struct {
	g.Meta `mime:"application/json" example:""`
}
