// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SuperAdmin is the golang structure of table super_admin for DAO operations like Where/Data.
type SuperAdmin struct {
	g.Meta       `orm:"table:super_admin, do:true"`
	Id           interface{} // 管理员ID
	Username     interface{} // 用户名
	PasswordHash interface{} // 密码
	Salt         interface{} // 密码盐
	Status       interface{} // 状态 normal:启用 disabled:禁用
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
