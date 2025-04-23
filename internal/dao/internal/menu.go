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
	Id          string // 菜单ID
	Pid         string // 父菜单ID
	Level       string // 关系树等级
	Tree        string // 关系树
	Title       string // 菜单名称
	Name        string // 菜单标识
	Path        string // 路由地址
	Component   string // 组件路径
	Type        string // 菜单类型 directory:目录 menu:菜单 button:按钮
	Permissions string // 权限标识
	Status      string // 状态 normal:启用 disabled:禁用
	Sort        string // 排序
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// menuColumns holds the columns for the table menu.
var menuColumns = MenuColumns{
	Id:          "id",
	Pid:         "pid",
	Level:       "level",
	Tree:        "tree",
	Title:       "title",
	Name:        "name",
	Path:        "path",
	Component:   "component",
	Type:        "type",
	Permissions: "permissions",
	Status:      "status",
	Sort:        "sort",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
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
