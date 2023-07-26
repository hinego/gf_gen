
package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

type InfoReq struct {
	g.Meta `path:"/v2/data/user/info" tags:"User" method:"post" summary:""`
}
type InfoRes struct {
	g.Meta `mime:"application/json" example:""`
}
