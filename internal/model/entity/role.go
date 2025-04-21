// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint64      `json:"id"        orm:"id"         description:"角色ID"`         // 角色ID
	Name      string      `json:"name"      orm:"name"       description:"角色名称"`         // 角色名称
	Key       string      `json:"key"       orm:"key"        description:"角色权限字符串"`      // 角色权限字符串
	Status    bool        `json:"status"    orm:"status"     description:"状态 1:启用 0:禁用"` // 状态 1:启用 0:禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`         // 删除时间
}
