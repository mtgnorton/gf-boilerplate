package role

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-boilerplate/internal/consts/enum"
	"gf-boilerplate/internal/variable"
)

type RoleArg struct {
	Name   string      `json:"name"   v:"required|length:3,10|unique-field:role#cg.role.name.required|cg.role.name.length|cg.role.name.unique" example:"管理员"    dc:"角色名称"`
	Code   string      `json:"code"   v:"required|length:3,10|unique-field:role#cg.role.code.required|cg.role.code.length|cg.role.code.unique" example:"admin"  dc:"角色权限字符串"`
	Status enum.Status `json:"status" v:"required|enums#cg.role.status.required|cg.role.status.enums"                                          example:"normal" dc:"状态,disabled:禁用,normal:正常"`
}

type RoleRet struct {
	Id             int64 `json:"id"      example:"1" dc:"角色ID"`
	RoleArg        `      json:",inline"             dc:"角色"`
	variable.Times `      json:",inline"             dc:"时间"`
}

type CreateReq struct {
	g.Meta `path:"/role/create" tags:"权限/角色" method:"post" summary:"创建角色"`
	RoleArg
}

type CreateRes struct {
	Id int64 `json:"id" dc:"角色ID"`
}

type UpdateReq struct {
	g.Meta `      path:"/role/update" tags:"权限/角色" method:"put" summary:"更新角色"`
	Id     int64 `                                                             json:"id" v:"required|exist-record:role#cg.role.id.required|cg.role.not_found" dc:"角色ID"`
	RoleArg
}

type UpdateRes struct {
	Id int64 `json:"id" dc:"角色ID"`
}

type DeleteReq struct {
	g.Meta `      path:"/role/delete" tags:"权限/角色" method:"delete" summary:"删除角色"`
	Id     int64 `                                                                v:"required" dc:"角色ID"`
}

type DeleteRes struct {
	Id int64 `json:"id" dc:"角色ID"`
}

type GetOneReq struct {
	g.Meta `      path:"/role/one" tags:"权限/角色" method:"get" summary:"获取角色详情"`
	Id     int64 `                                                            v:"required" example:"1" dc:"角色ID"`
}

type GetOneRes struct {
	Id int64 `json:"id" example:"1" dc:"角色ID"`
	RoleArg
}

type GetListReq struct {
	g.Meta `path:"/role/list" tags:"权限/角色" method:"get" summary:"获取角色列表"`
}
type GetListRes struct {
	List []RoleRet `json:"list" dc:"角色列表"`
}

type GetPageListReq struct {
	g.Meta `path:"/role/page_list" tags:"权限/角色" method:"get" summary:"获取角色分页列表"`
	variable.PaginationReq
}

type GetPageListRes struct {
	List                   []RoleRet `json:"list"    dc:"角色列表"`
	variable.PaginationRes `          json:",inline" dc:"分页"`
}
