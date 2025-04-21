// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleMenuMapping is the golang structure of table role_menu_mapping for DAO operations like Where/Data.
type RoleMenuMapping struct {
	g.Meta    `orm:"table:role_menu_mapping, do:true"`
	Id        interface{} // ID
	RoleId    interface{} // 角色ID
	MenuId    interface{} // 菜单ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
