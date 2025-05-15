// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SuperAdminDao is the data access object for the table super_admin.
type SuperAdminDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SuperAdminColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SuperAdminColumns defines and stores column names for the table super_admin.
type SuperAdminColumns struct {
	Id           string // 管理员ID
	Username     string // 用户名
	PasswordHash string // 密码
	Salt         string // 密码盐
	Status       string // 状态 normal:启用 disabled:禁用
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// superAdminColumns holds the columns for the table super_admin.
var superAdminColumns = SuperAdminColumns{
	Id:           "id",
	Username:     "username",
	PasswordHash: "password_hash",
	Salt:         "salt",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewSuperAdminDao creates and returns a new DAO object for table data access.
func NewSuperAdminDao(handlers ...gdb.ModelHandler) *SuperAdminDao {
	return &SuperAdminDao{
		group:    "default",
		table:    "super_admin",
		columns:  superAdminColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SuperAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SuperAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SuperAdminDao) Columns() SuperAdminColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SuperAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SuperAdminDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SuperAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
