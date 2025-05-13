// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for the table menu.
type MenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MenuColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MenuColumns defines and stores column names for the table menu.
type MenuColumns struct {
	Id            string // 菜单ID
	Pid           string // 父菜单ID,pid=0时为顶级菜单
	Level         string // 关系树等级,pid=0的menu等级为1,然后根据父子关系依次递增
	Tree          string // 关系树,将当前menu的所有父级menu的id用逗号分隔拼接而成,顺序从直接上级到最顶级
	Name          string // 菜单名称
	Code          string // 菜单标识
	Path          string // 前端路由地址
	Component     string // 前端组件路径
	Type          string // 菜单类型 directory:目录 menu:菜单 button:按钮
	Permissions   string // 该menu对应的后端路由(权限标识),如果一个菜单页面具有多个后端路由,用逗号分隔
	Icon          string // 菜单图标
	Redirect      string // 重定向地址
	BarActiveCode string // 当前页面高亮哪个code对应的左侧菜单,例如,在/user/edit页面时,左侧菜单仍然高亮/user/list
	IsExternal    string // 是否为外链,0:否 1:是
	ExternalUrl   string // 外链地址
	Status        string // 状态 normal:启用 disabled:禁用
	Sort          string // 排序
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// menuColumns holds the columns for the table menu.
var menuColumns = MenuColumns{
	Id:            "id",
	Pid:           "pid",
	Level:         "level",
	Tree:          "tree",
	Name:          "name",
	Code:          "code",
	Path:          "path",
	Component:     "component",
	Type:          "type",
	Permissions:   "permissions",
	Icon:          "icon",
	Redirect:      "redirect",
	BarActiveCode: "bar_active_code",
	IsExternal:    "is_external",
	ExternalUrl:   "external_url",
	Status:        "status",
	Sort:          "sort",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao(handlers ...gdb.ModelHandler) *MenuDao {
	return &MenuDao{
		group:    "default",
		table:    "menu",
		columns:  menuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
