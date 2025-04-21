// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Casbin is the golang structure for table casbin.
type Casbin struct {
	Id        uint64      `json:"id"        orm:"id"         description:"ID"`   // ID
	PType     string      `json:"pType"     orm:"p_type"     description:"策略类型"` // 策略类型
	V0        string      `json:"v0"        orm:"v0"         description:"主体"`   // 主体
	V1        string      `json:"v1"        orm:"v1"         description:"资源"`   // 资源
	V2        string      `json:"v2"        orm:"v2"         description:"操作"`   // 操作
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"` // 删除时间
}
