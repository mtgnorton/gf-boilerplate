// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CasbinDao is the data access object for the table casbin.
type CasbinDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CasbinColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CasbinColumns defines and stores column names for the table casbin.
type CasbinColumns struct {
	Id        string // ID
	PType     string // 策略类型
	V0        string // 主体
	V1        string // 资源
	V2        string // 操作
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// casbinColumns holds the columns for the table casbin.
var casbinColumns = CasbinColumns{
	Id:        "id",
	PType:     "p_type",
	V0:        "v0",
	V1:        "v1",
	V2:        "v2",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCasbinDao creates and returns a new DAO object for table data access.
func NewCasbinDao(handlers ...gdb.ModelHandler) *CasbinDao {
	return &CasbinDao{
		group:    "default",
		table:    "casbin",
		columns:  casbinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CasbinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CasbinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CasbinDao) Columns() CasbinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CasbinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CasbinDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CasbinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
