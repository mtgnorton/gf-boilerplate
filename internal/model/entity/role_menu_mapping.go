// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleMenuMapping is the golang structure for table role_menu_mapping.
type RoleMenuMapping struct {
	Id        uint64      `json:"id"        orm:"id"         description:"ID"`   // ID
	RoleId    uint64      `json:"roleId"    orm:"role_id"    description:"角色ID"` // 角色ID
	MenuId    uint64      `json:"menuId"    orm:"menu_id"    description:"菜单ID"` // 菜单ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"` // 删除时间
}
