// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id            uint64      `json:"id"            orm:"id"              ` // 菜单ID
	Pid           uint64      `json:"pid"           orm:"pid"             ` // 父菜单ID,pid=0时为顶级菜单
	Level         uint        `json:"level"         orm:"level"           ` // 关系树等级,pid=0的menu等级为1,然后根据父子关系依次递增
	Tree          string      `json:"tree"          orm:"tree"            ` // 关系树,将当前menu的所有父级menu的id用逗号分隔拼接而成,顺序从直接上级到最顶级
	Name          string      `json:"name"          orm:"name"            ` // 菜单名称
	Code          string      `json:"code"          orm:"code"            ` // 菜单标识
	Path          string      `json:"path"          orm:"path"            ` // 前端路由地址
	Component     string      `json:"component"     orm:"component"       ` // 前端组件路径
	Type          string      `json:"type"          orm:"type"            ` // 菜单类型 directory:目录 menu:菜单 button:按钮
	Permissions   string      `json:"permissions"   orm:"permissions"     ` // 该menu对应的后端路由(权限标识),如果一个菜单页面具有多个后端路由,用逗号分隔
	Icon          string      `json:"icon"          orm:"icon"            ` // 菜单图标
	Redirect      string      `json:"redirect"      orm:"redirect"        ` // 重定向地址
	BarActiveCode string      `json:"barActiveCode" orm:"bar_active_code" ` // 当前页面高亮哪个code对应的左侧菜单,例如,在/user/edit页面时,左侧菜单仍然高亮/user/list
	IsExternal    bool        `json:"isExternal"    orm:"is_external"     ` // 是否为外链,0:否 1:是
	ExternalUrl   string      `json:"externalUrl"   orm:"external_url"    ` // 外链地址
	Status        string      `json:"status"        orm:"status"          ` // 状态 normal:启用 disabled:禁用
	Sort          uint        `json:"sort"          orm:"sort"            ` // 排序
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      ` // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      ` // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      ` // 删除时间
}
