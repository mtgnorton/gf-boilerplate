package role

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta `path:"/role" tags:"角色" method:"post" summary:"创建角色"`
	Name   string `json:"name" v:"required#角色名称不能为空"`
	Code   string `json:"code" v:"required#角色编码不能为空"`
}

type CreateRes struct {
	Id string `json:"id"`
}
