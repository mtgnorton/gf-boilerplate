package member

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-boilerplate/internal/consts/enum"
)

type Member struct {
	RoleId            uint64      `v:"required"               dc:"角色ID"`
	Username          string      `v:"required|length:6,10"   dc:"用户名"`
	Password          string      `v:"required|password2"     dc:"密码"`
	PasswordConfirmed string      `v:"required|same:Password" dc:"确认密码"`
	Status            enum.Status `v:"required|enums"         dc:"状态,normal:正常,disabled:禁用"`
}

type CreateReq struct {
	g.Meta `path:"/member" tags:"权限/用户" method:"post" summary:"创建用户"`
	Member
}

type CreateRes struct {
	Id string `json:"id" dc:"用户ID"`
}

type UpdateReq struct {
	g.Meta `       path:"/member/:id" tags:"权限/用户" method:"put" summary:"更新用户"`
	Id     string `                                                            v:"required" dc:"用户ID"`
	Member
}

type UpdateRes struct {
	Id string `json:"id" dc:"用户ID"`
}
