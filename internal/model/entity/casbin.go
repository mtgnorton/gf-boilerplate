// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Casbin is the golang structure for table casbin.
type Casbin struct {
	Id        uint64      `json:"id"        orm:"id"         ` // ID
	PType     string      `json:"pType"     orm:"p_type"     ` // 策略类型
	V0        string      `json:"v0"        orm:"v0"         ` // 主体
	V1        string      `json:"v1"        orm:"v1"         ` // 资源
	V2        string      `json:"v2"        orm:"v2"         ` // 操作
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // 删除时间
}
