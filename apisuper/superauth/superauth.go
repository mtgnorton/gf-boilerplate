package superauth

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-boilerplate/internal/variable"
)

// 管理员参数
type MemberArg struct {
	RoleId   int64  `json:"role_id"  v:"required#cg.member.role_id.required"                                                                                      example:"1"      dc:"角色ID"`
	Username string `json:"username" v:"required|length:2,32|unique-field:member#cg.member.username.required|cg.member.username.length|cg.member.username.unique" example:"admin"  dc:"用户名"`
	Password string `json:"password" v:"required-without:id|length:6,32#cg.member.password.required|cg.member.password.length"                                    example:"123456" dc:"密码"`
	Status   string `json:"status"   v:"required|in:normal,disabled#cg.member.status.required|cg.member.status.in"                                                example:"normal" dc:"状态,normal:启用,disabled:禁用"`
}

// 管理员返回
type MemberRet struct {
	Id       int64  `json:"id"       example:"1"      dc:"管理员ID"`
	RoleId   int64  `json:"role_id"  example:"1"      dc:"角色ID"`
	Username string `json:"username" example:"admin"  dc:"用户名"`
	Status   string `json:"status"   example:"normal" dc:"状态"`
	variable.Times
}

// 创建管理员
type CreateReq struct {
	g.Meta `path:"/member/create" tags:"权限/管理员" method:"post" summary:"创建管理员"`
	MemberArg
}
type CreateRes struct {
	Id int64 `json:"id" dc:"管理员ID"`
}

// 更新管理员
type UpdateReq struct {
	g.Meta `      path:"/member/update" tags:"权限/管理员" method:"post" summary:"更新管理员"`
	Id     int64 `                                                                  json:"id" v:"required#cg.member.id.required" dc:"管理员ID"`
	MemberArg
}
type UpdateRes struct {
	Id int64 `json:"id" dc:"管理员ID"`
}

// 删除管理员
type DeleteReq struct {
	g.Meta `      path:"/member/delete" tags:"权限/管理员" method:"post" summary:"删除管理员"`
	Id     int64 `                                                                  json:"id" v:"required#cg.member.id.required" dc:"管理员ID"`
}
type DeleteRes struct {
	Id int64 `json:"id" dc:"管理员ID"`
}

// 获取管理员详情
type GetOneReq struct {
	g.Meta `      path:"/member/one" tags:"权限/管理员" method:"get" summary:"获取管理员详情"`
	Id     int64 `                                                                json:"id" v:"required#cg.member.id.required" example:"1" dc:"管理员ID"`
}
type GetOneRes struct {
	Id int64 `json:"id" example:"1" dc:"管理员ID"`
	MemberRet
}

// 获取管理员列表
type GetListReq struct {
	g.Meta `path:"/member/list" tags:"权限/管理员" method:"get" summary:"获取管理员列表"`
}
type GetListRes struct {
	List []MemberRet `json:"list" dc:"管理员列表"`
}

// 获取管理员分页列表
type GetPageListReq struct {
	g.Meta `path:"/member/page_list" tags:"权限/管理员" method:"get" summary:"获取管理员分页列表"`
	variable.PaginationReq
}
type GetPageListRes struct {
	List                   []MemberRet `json:"list"    dc:"管理员列表"`
	variable.PaginationRes `            json:",inline" dc:"分页"`
}
