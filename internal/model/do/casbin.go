// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Casbin is the golang structure of table casbin for DAO operations like Where/Data.
type Casbin struct {
	g.Meta    `orm:"table:casbin, do:true"`
	Id        interface{} // ID
	PType     interface{} // 策略类型
	V0        interface{} // 主体
	V1        interface{} // 资源
	V2        interface{} // 操作
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
