// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleMenuMappingDao is the data access object for the table role_menu_mapping.
type RoleMenuMappingDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  RoleMenuMappingColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// RoleMenuMappingColumns defines and stores column names for the table role_menu_mapping.
type RoleMenuMappingColumns struct {
	Id        string // ID
	RoleId    string // 角色ID
	MenuId    string // 菜单ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// roleMenuMappingColumns holds the columns for the table role_menu_mapping.
var roleMenuMappingColumns = RoleMenuMappingColumns{
	Id:        "id",
	RoleId:    "role_id",
	MenuId:    "menu_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewRoleMenuMappingDao creates and returns a new DAO object for table data access.
func NewRoleMenuMappingDao(handlers ...gdb.ModelHandler) *RoleMenuMappingDao {
	return &RoleMenuMappingDao{
		group:    "default",
		table:    "role_menu_mapping",
		columns:  roleMenuMappingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RoleMenuMappingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RoleMenuMappingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RoleMenuMappingDao) Columns() RoleMenuMappingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RoleMenuMappingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RoleMenuMappingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RoleMenuMappingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
