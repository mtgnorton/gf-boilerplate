// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure for table member.
type Member struct {
	Id        uint64      `json:"id"        orm:"id"         description:"管理员ID"`                    // 管理员ID
	RoleId    uint64      `json:"roleId"    orm:"role_id"    description:"角色ID"`                     // 角色ID
	Username  string      `json:"username"  orm:"username"   description:"用户名"`                      // 用户名
	Password  string      `json:"password"  orm:"password"   description:"密码"`                       // 密码
	Salt      string      `json:"salt"      orm:"salt"       description:"密码盐"`                      // 密码盐
	Status    string      `json:"status"    orm:"status"     description:"状态 normal:启用 disabled:禁用"` // 状态 normal:启用 disabled:禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                     // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                     // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`                     // 删除时间
}
