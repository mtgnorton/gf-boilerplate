package menu

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-boilerplate/internal/consts/enum"
	"gf-boilerplate/internal/variable"
)

type Arg struct {
	Pid           uint64        `json:"pid"           example:"1"                      dc:"父菜单ID"`
	Name          string        `json:"name"          example:"菜单名称"                   dc:"菜单名称"                                                     v:"required|length:2,50#cg.menu.name.required|cg.menu.name.length"`
	Code          string        `json:"code"          example:"menu"                   dc:"菜单标识"                                                     v:"required|length:4,50|unique-field:menu#cg.menu.code.required|cg.menu.code.length|cg.menu.code.unique"`
	Path          string        `json:"path"          example:"/menu"                  dc:"前端路由地址"                                                   v:"required|length:2,100#cg.menu.path.required|cg.menu.path.length"`
	Component     string        `json:"component"     example:"/menu"                  dc:"前端组件路径"                                                   v:"required|length:2,255#cg.menu.component.required|cg.menu.component.length"`
	Type          enum.MenuType `json:"type"          example:"directory"              dc:"菜单类型,directory:目录 menu:菜单 button:按钮"                      v:"required|enums#cg.menu.type.required|cg.menu.type.enums"`
	Permissions   string        `json:"permissions"   example:"menu:list,menu:create"  dc:"该menu对应的后端路由(权限标识),如果一个菜单页面具有多个后端路由,用逗号分隔"                v:"length:4,255#cg.menu.permissions.length"`
	Icon          string        `json:"icon"          example:""                       dc:"菜单图标"                                                     v:"max-length:255#cg.menu.icon.max-length"`
	Redirect      string        `json:"redirect"      example:"/menu"                  dc:"重定向地址"                                                    v:"max-length:255#cg.menu.redirect.max-length"`
	BarActiveCode string        `json:"barActiveCode" example:"menu:list"              dc:"当前页面高亮哪个code对应的左侧菜单,例如,在/user/edit页面时,左侧菜单仍然高亮/user/list" v:"max-length:50#cg.menu.barActiveCode.max-length"`
	IsExternal    bool          `json:"isExternal"    example:"false"                  dc:"是否为外链,0:否 1:是"`
	ExternalUrl   string        `json:"externalUrl"   example:"https://www.google.com" dc:"外链地址"                                                     v:"max-length:255#cg.menu.externalUrl.max-length"`
	Status        enum.Status   `json:"status"        example:"normal"                 dc:"状态,normal:启用 disabled:禁用"                                 v:"required|enums#cg.menu.status.required|cg.menu.status.enums"`
	Sort          uint          `json:"sort"          example:"1"                      dc:"排序"`
}

type TreeItem struct {
	Id             int64 `json:"id"                 example:"1" dc:"菜单ID"`
	Arg            `            json:",inline"                        dc:"菜单"`
	variable.Times `            json:",inline"                        dc:"时间"`
	Children       []*TreeItem `json:"children,omitempty"             dc:"子菜单"`
}

type CreateReq struct {
	g.Meta `path:"/menu/create" tags:"权限/菜单" method:"post" summary:"创建菜单"`
	Arg
}

type CreateRes struct {
	Id int64 `json:"id" dc:"菜单ID"`
}

type UpdateReq struct {
	g.Meta `      path:"/menu/update" tags:"权限/菜单" method:"post" summary:"更新菜单"`
	Id     int64 `                                                              json:"id" v:"required|exist-record:menu#cg.menu.id.required|cg.menu.not_found" dc:"菜单ID"`
	Arg
}

type UpdateRes struct {
	Id int64 `json:"id" dc:"菜单ID"`
}

type DeleteReq struct {
	g.Meta `      path:"/menu/delete" tags:"权限/菜单" method:"post" summary:"删除菜单"`
	Id     int64 `                                                              v:"required|exist-record:menu#cg.menu.id.required|cg.menu.not_found" dc:"菜单ID"`
}

type DeleteRes struct {
	Id int64 `json:"id" dc:"菜单ID"`
}

type GetListReq struct {
	g.Meta `path:"/menu/list" tags:"权限/菜单" method:"get" summary:"获取菜单列表"`
}

type GetListRes struct {
	List []*TreeItem `json:"list" dc:"菜单列表"`
}
