// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id          uint64      `json:"id"          orm:"id"          description:"菜单ID"`                                // 菜单ID
	Pid         uint64      `json:"pid"         orm:"pid"         description:"父菜单ID"`                               // 父菜单ID
	Level       uint        `json:"level"       orm:"level"       description:"关系树等级"`                               // 关系树等级
	Tree        string      `json:"tree"        orm:"tree"        description:"关系树"`                                 // 关系树
	Title       string      `json:"title"       orm:"title"       description:"菜单名称"`                                // 菜单名称
	Name        string      `json:"name"        orm:"name"        description:"菜单标识"`                                // 菜单标识
	Path        string      `json:"path"        orm:"path"        description:"路由地址"`                                // 路由地址
	Component   string      `json:"component"   orm:"component"   description:"组件路径"`                                // 组件路径
	Type        string      `json:"type"        orm:"type"        description:"菜单类型 directory:目录 menu:菜单 button:按钮"` // 菜单类型 directory:目录 menu:菜单 button:按钮
	Permissions string      `json:"permissions" orm:"permissions" description:"权限标识"`                                // 权限标识
	Status      bool        `json:"status"      orm:"status"      description:"状态 1:启用 0:禁用"`                        // 状态 1:启用 0:禁用
	Sort        uint        `json:"sort"        orm:"sort"        description:"排序"`                                  // 排序
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`                                // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`                                // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`                                // 删除时间
}
