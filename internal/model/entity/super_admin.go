// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SuperAdmin is the golang structure for table super_admin.
type SuperAdmin struct {
	Id           uint64      `json:"id"           orm:"id"            ` // 管理员ID
	Username     string      `json:"username"     orm:"username"      ` // 用户名
	PasswordHash string      `json:"passwordHash" orm:"password_hash" ` // 密码
	Salt         string      `json:"salt"         orm:"salt"          ` // 密码盐
	Status       string      `json:"status"       orm:"status"        ` // 状态 normal:启用 disabled:禁用
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    ` // 删除时间
}
