package access

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `       path:"/access/login" tags:"权限/登录" method:"post" summary:"登录"`
	Username string `                                                             json:"username" v:"required|length:4,20#cg.access.username.required|cg.access.username.length" example:"admin"  dc:"用户名"`
	Password string `                                                             json:"password" v:"required|length:6,20#cg.access.password.required|cg.access.password.length" example:"123456" dc:"密码"`
}

type LoginRes struct {
	Token string `json:"token" dc:"访问令牌"`
}

type LogoutReq struct {
	g.Meta `path:"/access/logout" tags:"权限/登录" method:"post" summary:"退出登录"`
}

type LogoutRes struct {
	Success bool `json:"success" dc:"是否成功"`
}
