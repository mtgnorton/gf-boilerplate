// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint64      `json:"id"        orm:"id"         ` // 角色ID
	Name      string      `json:"name"      orm:"name"       ` // 角色名称
	Code      string      `json:"code"      orm:"code"       ` // 角色权限字符串
	Status    string      `json:"status"    orm:"status"     ` // 状态 normal:启用 disabled:禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // 删除时间
}
