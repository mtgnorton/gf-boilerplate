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
	g.Meta        `orm:"table:menu, do:true"`
	Id            interface{} // 菜单ID
	Pid           interface{} // 父菜单ID,pid=0时为顶级菜单
	Level         interface{} // 关系树等级,pid=0的menu等级为1,然后根据父子关系依次递增
	Tree          interface{} // 关系树,将当前menu的所有父级menu的id用逗号分隔拼接而成,顺序从直接上级到最顶级
	Name          interface{} // 菜单名称
	Code          interface{} // 菜单标识
	Path          interface{} // 前端路由地址
	Component     interface{} // 前端组件路径
	Type          interface{} // 菜单类型 directory:目录 menu:菜单 button:按钮
	Permissions   interface{} // 该menu对应的后端路由(权限标识),如果一个菜单页面具有多个后端路由,用逗号分隔
	Icon          interface{} // 菜单图标
	Redirect      interface{} // 重定向地址
	BarActiveCode interface{} // 当前页面高亮哪个code对应的左侧菜单,例如,在/user/edit页面时,左侧菜单仍然高亮/user/list
	IsExternal    interface{} // 是否为外链,0:否 1:是
	ExternalUrl   interface{} // 外链地址
	Status        interface{} // 状态 normal:启用 disabled:禁用
	Sort          interface{} // 排序
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
