// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta      `orm:"table:menu, do:true"`
	Id          interface{} // 菜单ID
	Pid         interface{} // 父菜单ID
	Level       interface{} // 关系树等级
	Tree        interface{} // 关系树
	Title       interface{} // 菜单名称
	Name        interface{} // 菜单标识
	Path        interface{} // 路由地址
	Component   interface{} // 组件路径
	Type        interface{} // 菜单类型 directory:目录 menu:菜单 button:按钮
	Permissions interface{} // 权限标识
	Status      interface{} // 状态 1:启用 0:禁用
	Sort        interface{} // 排序
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
